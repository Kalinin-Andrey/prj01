package main

import (
	"log"

	"carizza/internal/pkg/config"

	commonApp "carizza/internal/app"
	"carizza/internal/app/restapi"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatalln("Can not load the config")
	}
	app := restapi.New(commonApp.New(*cfg), *cfg)

	if err := app.Run(); err != nil {
		log.Fatalf("Error while application is running: %s", err.Error())
	}
	defer func() {
		if err := app.Stop(); err != nil {
			log.Fatalf("Error while application is stopping: %s", err.Error())
		}
	}()
}
