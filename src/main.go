package main

import (
	"log"
	"os"
	app "vmctl/src/app"
	config "vmctl/src/config"
)

func test() {

	config.LoadConfig(&config.Cfg)
	var app = app.NewApp(&config.Cfg)

	if err := app.Cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
