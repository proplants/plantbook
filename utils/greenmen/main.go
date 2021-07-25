package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/proplants/plantbook/pkg/logging"
	"github.com/proplants/plantbook/utils/greenmen/model"

	"github.com/gocolly/colly/v2"
)

const (
	letterBytes        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	version     string = "0.0.1"
	// nolint:lll
	templatePlants string = `INSERT INTO reference.plants (title, category_id, short_info, notes, img_links, created_at, creator) VALUES('%s','%s', '%s'::jsonb, '%s'::jsonb, '%s'::jsonb, CURRENT_TIMESTAMP, CURRENT_USER);`

	rndLengthLim int = 10

	expectedPlantsLength int = 310
	chURLLength          int = 0
)

var rPPlants = []model.RefPagePlants{
	{
		Name:       "roomPlants",
		URL:        "http://www.plantopedia.ru/encyclopaedia/pot-plant/sections.php",
		FileName:   "room_plants",
		CategoryID: "1",
	},
	{
		Name:       "gardenPlants",
		URL:        "http://www.plantopedia.ru/encyclopaedia/garden-plants/sections.php",
		FileName:   "garden_plants",
		CategoryID: "2",
	},
	{
		Name:       "cuttingPlants",
		URL:        "http://www.plantopedia.ru/encyclopaedia/cutting-plants/sections.php",
		FileName:   "cutting_plants",
		CategoryID: "3",
	},
	{
		Name:       "ogorodPlants",
		URL:        "http://www.plantopedia.ru/encyclopaedia/ogorod/sections.php",
		FileName:   "ogorod_plants",
		CategoryID: "4",
	},
}

// RandomString ...
// nolint:gosec
func RandomString() string {
	b := make([]byte, rand.Intn(rndLengthLim)+rndLengthLim)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	// flags
	debug := flag.Bool("d", false, "true/false/1/0 debug on/off")
	directLink := flag.String("u", "", "url to plant page for single scrapping")

	flag.Parse()

	// set logger
	logger := logging.NewLogger(*debug, "console")
	logger = logger.With("version", version)
	mainctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx := logging.WithLogger(mainctx, logger)

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("www.plantopedia.ru", "plantopedia.ru"),
	)
	cc := newCC(c)
	if *directLink != "" {
		singlePlant, err := cc.parsePlantPage(ctx, *directLink)
		if err != nil {
			logger.Errorf("parsePlantPage error, %s", err)
			return
		}
		fmt.Printf("plant: %s\n", singlePlant)
		return
	}
	for _, rp := range rPPlants {
		plants := make(model.Plants, 0, expectedPlantsLength)
		chURLs := make(chan string, chURLLength)
		done := make(chan bool)
		go func() {
			plants = cc.goParsePlants(ctx, chURLs, done)
		}()

		err := cc.parseRefPage(ctx, rp, chURLs)
		if err != nil {
			logger.Infof("cc.parseRefPage URL: %v error: , %s", rp.URL, err)
		}
		close(chURLs)
		<-done
		logger.Infof("parsed %d url(s) in %s", len(plants), rp.Name)
		err = saveToFile(plants, "./parsed/"+rp.FileName+".json", "./parsed/"+rp.FileName+".sql", rp.CategoryID)
		if err != nil {
			logger.Errorf("saveToFile error to %v: , %s", rp.FileName, err)
		}
	}
}

// helpers

func saveToFile(plants model.Plants, fnData, fnSQL, rpCategory string) error {
	f, err := os.Create(fnData)
	if err != nil {
		return err
	}
	defer f.Close()
	fq, err := os.Create(fnSQL)
	if err != nil {
		return err
	}
	defer fq.Close()
	for _, plant := range plants {
		// nolint:errcheck
		// fmt.Fprintln()
		fq.WriteString(makeSQLInsert(plant, rpCategory) + "\n")
	}
	data, err := plants.MarshalJSON()
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	return err
}

// nolint
func makeSQLInsert(plant model.Plant, category string) string {
	// INSERT INTO reference.plants (title, category_id, short_info, notes, img_links, created_at, creator)
	// VALUES('%s', %s, '%s'::jsonb, '%s'::jsonb, '%s'::jsonb, CURRENT_TIMESTAMP, CURRENT_USER);
	shortInfoBts, _ := plant.ShortInfo.MarshalJSON()
	notesBts, _ := plant.Info.MarshalJSON()
	imgLinksBts, _ := plant.Images.MarshalJSON()
	return fmt.Sprintf(templatePlants, plant.Title, category, shortInfoBts, notesBts, imgLinksBts)
}

func (c *Collector) goParsePlants(ctx context.Context, chURLs chan string, done chan bool) model.Plants {
	logger := logging.FromContext(ctx)
	plants := make(model.Plants, 0, expectedPlantsLength)
	for {
		select {
		case <-ctx.Done():
			done <- true
			logger.Debugf("ctx.Done")
			return plants
		case url, more := <-chURLs:
			if !more {
				done <- true
				logger.Debugf("channel chURLs close, more: %v", more)
				return plants
			}
			plant, err := c.parsePlantPage(ctx, url)
			if err != nil {
				logger.Errorf("parsePlantPage error %s", err)
				continue
			}
			plants = append(plants, *plant)
		}
	}
}
