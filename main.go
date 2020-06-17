package main

import (
	"gotrading/app/controllers"
	"log"
)

func main() {
	// utils.LoggingSettings(config.Config.LogFile)
	// controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())
}
