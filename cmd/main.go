package main

import (
	"avatg/internal/app/config"
	"avatg/internal/app/di"
	"fmt"
	"log"
)

func main() {
	config, err := config.ProvideConfig()
	fmt.Println(config)
	if err != nil {
		log.Fatal("app read settings: ", err)
	}

	app, cleanupResourcesFunc, err := di.InitializeApp(config)
	if err != nil {
		log.Fatal("app start: ", err)
	}

	app.Run(cleanupResourcesFunc)
}
