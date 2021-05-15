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

	log.Printf(
		"PORT: %v\n, HOST_DB: %v\n,PORT_DB: %v\n,  DB_USER: %v\n, DB_NAME: %v\n",
		 config.Elements.Port,
		 config.Elements.DBHost,
		 config.Elements.DBPort,
		 config.Elements.DBUser,
		 config.Elements.DBName) // TODO change logger in the future
}
