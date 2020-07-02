package controllers

import (
	"gotrading/app/models"
	"gotrading/bitflyer"
	"gotrading/config"
	"log"
)

// StreamIngestionData is データの取得とトレードの開始
func StreamIngestionData() {
	c := config.Config
	e := config.Env
	ai := NewAI(c.ProductCode, c.TradeDuration, c.DataLimit, c.UsePercent, c.StopLimitPercent, e.BackTest)

	var tickerChannl = make(chan bitflyer.Ticker)
	apiClient := bitflyer.New(e.ApiKey, e.ApiSecret)
	go apiClient.GetRealTimeTicker(c.ProductCode, tickerChannl)
	go func() {
		for ticker := range tickerChannl {
			log.Printf("action=StreamIngestionData, %v", ticker)
			for _, duration := range c.Durations {
				isCreated := models.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
				if isCreated == true && duration == c.TradeDuration {
					go ai.Trade()
				}
			}
		}
	}()
}
