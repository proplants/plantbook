package main

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/proplants/plantbook/pkg/logging"
	"github.com/proplants/plantbook/utils/greenmen/model"

	"github.com/gocolly/colly/v2"
	"github.com/pkg/errors"
)

// plant page
// [x] room plants

const (
	shortPropKind                string = "Тип растения"
	shortPropRecomendPosition    string = "Рекомендуемое расположение"
	shortPropReg2Light           string = "Отношение к свету"
	shortPropReg2Moisture        string = "Отношение к влаге"
	shortPropFloweringTime       string = "Сроки цветения"
	shortPropHight               string = "Высота"
	shortPropClassifiers         string = "Ценность в культуре"
	shortPropGround              string = "Почва"
	shortPropWintering           string = "Зимовка"
	shortPropDecorativeness      string = "Форма декоративности"
	shortPropComposition         string = "Значимость в композиции"
	shortPropShearing            string = "Устойчивость в срезке"
	shortPropGrowing             string = "Условия выращивания"
	shortPropEating              string = "Употребление в пищу"
	minPartsCountOfTheDivPlashka int    = 2
)

// Collector html grabber.
type Collector struct {
	c *colly.Collector
}

func newCC(c *colly.Collector) *Collector {
	return &Collector{c: c}
}

func (c *Collector) parseRefPage(ctx context.Context, rpp model.RefPagePlants, out chan string) error {
	cc := c.c.Clone()
	log := logging.FromContext(ctx)
	u, err := url.Parse(rpp.URL)
	if err != nil {
		return errors.WithMessage(err, "parse pageURL error")
	}
	baseURL := u.Scheme + "://" + u.Host
	var href string
	log.Debugf("got baseURL %s", baseURL)
	cc.OnHTML(".kolon", func(e *colly.HTMLElement) {
		e.DOM.Children().Children().Children().Children().Each(func(i int, ee *goquery.Selection) {
			if rpp.Name == "roomPlants" {
				href, _ = ee.ChildrenFiltered("p").Children().Attr("href") // for pot_plants
			} else {
				href, _ = ee.ChildrenFiltered("a").Attr("href") // for garden_plants, cutting_plants, ogorod_plants
			}
			if len(href) == 0 {
				return
			}
			outURL := baseURL + href
			out <- outURL
			log.Debugf("send out url %s", outURL)
		})
	})

	// Before making a request print "Visiting ..."
	cc.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
		fmt.Println("Visiting", r.URL.String())
	})

	err = cc.Visit(rpp.URL)
	if err != nil {
		return errors.WithMessagef(err, "visit %s error", rpp.URL)
	}
	return nil
}

func (c *Collector) parsePlantPage(ctx context.Context, pageURL string) (*model.Plant, error) {
	cc := c.c.Clone()
	log := logging.FromContext(ctx)
	u, err := url.Parse(pageURL)
	if err != nil {
		return nil, errors.WithMessage(err, "parse pageURL error")
	}
	baseURL := u.Scheme + "://" + u.Host
	log.Debugf("got baseURL %s", baseURL)

	p := &model.Plant{
		Source:   pageURL,
		Title:    "",
		Category: "",
		ShortInfo: model.ShortInfo{
			Kind:              "",
			RecommendPosition: "",
			RegardToLight:     "",
			RegardToMoisture:  "",
			FloweringTime:     "",
			Hight:             "",
			Classifiers:       "",
			Ground:            "",
			Wintering:         "",
			Decorativeness:    "",
			Composition:       "",
			Shearing:          "",
			Growing:           "",
			Eating:            "",
		},
		Images: make([]string, 0, 4),
		Info:   make([]model.Info, 0, 4),
		Metadata: model.Metadata{
			DateCollect: time.Now().Format(time.RFC3339),
			Target:      pageURL,
		},
	}
	// set title
	cc.OnHTML(".encyclopaedia-info", func(e *colly.HTMLElement) {
		// category
		if category, ok := e.DOM.ChildrenFiltered("meta").Attr("content"); ok {
			p.Category = category
		}
		// title
		p.Title = e.DOM.ChildrenFiltered("h2").Text()
		log.Debugf("set category %s, and title %s", p.Category, p.Title)
	})

	// set short_info
	p.ShortInfo = setShortInfo(ctx, cc, &p.ShortInfo)
	log.Debugf("set shorInfo: %+v", &p.ShortInfo)

	// images
	cc.OnHTML("#pikame", func(e *colly.HTMLElement) {
		e.ForEach("img", func(i int, ee *colly.HTMLElement) {
			p.Images = append(p.Images, baseURL+ee.Attr("src"))
		})
		log.Debugf("set images: %v", p.Images)
	})

	const minLengthDivInfo int = 2
	// info
	cc.OnHTML(".encyclopaedia-zag", func(e *colly.HTMLElement) {
		e.ForEach("h3", func(i int, ee *colly.HTMLElement) {
			title := strings.Join(strings.Fields(strings.TrimSpace(ee.Text)), " ")
			content := strings.Join(strings.Fields(strings.TrimSpace(ee.DOM.Next().Text())), " ")
			log.Debugf("info title: %s, content: %s", title, content)
			if len(content) < minLengthDivInfo {
				content = strings.Join(strings.Fields(strings.TrimSpace(ee.DOM.Next().Next().Text())), " ")
				if len(content) < minLengthDivInfo {
					return
				}
			}
			content = strings.ReplaceAll(content, "'", "")
			p.Info = append(p.Info, model.Info{Title: title, Content: content})
		})
		log.Debugf("set info length: %d", len(p.Info))
	})

	// Before making a request print "Visiting ..."
	cc.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
		fmt.Println("Visiting", r.URL.String())
	})

	err = cc.Visit(pageURL)
	if err != nil {
		return p, errors.WithMessagef(err, "visit %s error", pageURL)
	}
	return p, nil
}

func setShortInfo(ctx context.Context, cc *colly.Collector, shortInfo *model.ShortInfo) model.ShortInfo {
	log := logging.FromContext(ctx)
	cc.OnHTML(".plashka", func(e *colly.HTMLElement) {
		e.ForEach("div", func(i int, ee *colly.HTMLElement) {
			keyValue := strings.Split(ee.Text, "\n")
			if len(keyValue) >= minPartsCountOfTheDivPlashka {
				prop := strings.TrimRight(strings.Join(strings.Fields(keyValue[0]), " "), ":")
				value := strings.Join(strings.Fields(strings.TrimSpace(keyValue[1])), " ")
				log.Debugf("prop: %s - value: %s", prop, value)
				switch prop {
				case shortPropKind:
					shortInfo.Kind = value
				case shortPropRecomendPosition:
					shortInfo.RecommendPosition = value
				case shortPropReg2Light:
					shortInfo.RegardToLight = value
				case shortPropReg2Moisture:
					shortInfo.RegardToMoisture = value
				case shortPropFloweringTime:
					shortInfo.FloweringTime = value
				case shortPropHight:
					shortInfo.Hight = value
				case shortPropClassifiers:
					shortInfo.Classifiers = value
				case shortPropGround:
					shortInfo.Ground = value
				case shortPropWintering:
					shortInfo.Wintering = value
				case shortPropDecorativeness:
					shortInfo.Decorativeness = value
				case shortPropComposition:
					shortInfo.Composition = value
				case shortPropShearing:
					shortInfo.Shearing = value
				case shortPropGrowing:
					shortInfo.Growing = value
				case shortPropEating:
					shortInfo.Eating = value
				}
			}
		})
	})
	log.Debugf("setShortInfo return: %+v", shortInfo)
	return *shortInfo
}
