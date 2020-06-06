package main

import (
	"gotrading/app/controllers"
	"gotrading/config"
	"gotrading/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())
}
