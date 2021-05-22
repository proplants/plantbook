package main

import (
	"log"

	"github.com/kaatinga/env_loader"
	"github.com/kaatinga/plantbook/config"
)

func main() {
	elements := config.New()
	// Environment variable initialization
	err := env_loader.LoadUsingReflect(&elements)
	if err != nil {
		log.Fatal(err) // TODO change logger in the future
	}
	config.ConfLog(elements)
}