package main

import (
	"log"
	"os"

	"carizza/internal/pkg/config"

	commonApp "carizza/internal/app"
	"carizza/internal/app/cli"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatalln("Can not load the config")
	}
	app := cli.New(commonApp.New(*cfg), *cfg)

	if err := app.Run(); err != nil {
		log.Fatalf("Error while cli application is running: %s", err.Error())
		os.Exit(1)
	}
}
