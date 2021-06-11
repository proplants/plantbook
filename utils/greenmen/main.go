package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/kaatinga/plantbook/pkg/logging"
	"github.com/kaatinga/plantbook/utils/greenmen/model"

	"github.com/gocolly/colly/v2"
)

const (
	letterBytes        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	version     string = "0.0.1"

	// nolint:lll
	templateRoomPlants string = `INSERT INTO reference.plants (title, category_id, short_info, notes, img_links, created_at, creator) VALUES('%s', 1, '%s'::jsonb, '%s'::jsonb, '%s'::jsonb, CURRENT_TIMESTAMP, CURRENT_USER);`
	rndLengthLim       int    = 10
)

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
	const (
		expectedPlantsLength int = 310
		chURLLength          int = 5
	)
	plants := make(model.Plants, 0, expectedPlantsLength)
	chURLs := make(chan string, chURLLength)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case url, more := <-chURLs:
				if !more {
					return
				}
				logger.Debugf("got url=%s", url)
				plant, err := cc.parsePlantPage(ctx, url)
				if err != nil {
					logger.Errorf("parsePlantPage error %s", err)
					continue
				}
				plants = append(plants, *plant)
			}
		}
	}()

	const refPageRoomPlants string = "http://www.plantopedia.ru/encyclopaedia/pot-plant/sections.php"

	err := cc.parseRefPage(ctx, refPageRoomPlants, chURLs)
	if err != nil {
		logger.Errorf("cc.parseRefPage error, %s", err)
	}
	logger.Infof("parsed %d url(s)", len(plants))
	err = saveToFile(plants, "plants.json", "room_plants.sql")
	if err != nil {
		logger.Errorf("saveToFile error, %s", err)
	}
}

// helpers

func saveToFile(plants model.Plants, fnData, fnSQL string) error {
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
		fq.WriteString(makeSQLInsert(plant, templateRoomPlants) + "\n")
	}
	data, err := plants.MarshalJSON()
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	return err
}

// nolint
func makeSQLInsert(plant model.Plant, template string) string {
	// INSERT INTO reference.plants (title, category_id, short_info, notes, img_links, created_at, creator)
	// VALUES('%s', 1, '%s'::jsonb, '%s'::jsonb, '%s'::jsonb, CURRENT_TIMESTAMP, CURRENT_USER);
	shortInfoBts, _ := plant.ShortInfo.MarshalJSON()
	notesBts, _ := plant.Info.MarshalJSON()
	imgLinksBts, _ := plant.Images.MarshalJSON()
	return fmt.Sprintf(template, plant.Title, shortInfoBts, notesBts, imgLinksBts)
}
