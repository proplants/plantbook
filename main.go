package main

import (
	"log"

	"github.com/kaatinga/env_loader"
	"github.com/kaatinga/plantbook/config"
)

func main() {
	// Environment variable initialization
	err := env_loader.LoadUsingReflect(&config.Elements)
	if err != nil {
		log.Fatal(err) // TODO change logger in the future
	}
	config.ConfLog()
}