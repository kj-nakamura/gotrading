package main

import (
	"errors"
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

func example(code string) (int, error) {
	if code == "hoge" {
		return 1, nil
	}
	return 0, errors.New("code must be hoge")
}
