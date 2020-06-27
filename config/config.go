package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type EnvValue struct {
	ApiKey      string `required:"true" split_words:"true"`
	ApiSecret   string `required:"true" split_words:"true"`
	DbName      string `required:"true" split_words:"true"`
	DbHost      string `required:"true" split_words:"true" default:"mysql"`
	DbUserName  string `required:"true" split_words:"true"`
	DbPassword  string `required:"true" split_words:"true"`
	IncomingURL string `split_words:"true"`
}

type ConfigValue struct {
	LogFile          string        `required:"true" split_words:"true"`
	ProductCode      string        `required:"true" split_words:"true"`
	TradeDuration    time.Duration `required:"true" split_words:"true"`
	BackTest         bool          `required:"true" split_words:"true"`
	UsePercent       float64       `required:"true" split_words:"true"`
	DataLimit        int           `required:"true" split_words:"true"`
	StopLimitPercent float64       `required:"true" split_words:"true"`
	NumRanking       int           `required:"true" split_words:"true"`
	Deadline         int           `required:"true" split_words:"true"`
	Durations        map[string]time.Duration
	SQLDriver        string `required:"true" split_words:"true"`
	Port             int    `required:"true" split_words:"true"`
}

var Env EnvValue
var Config ConfigValue

func init() {
	if err := envconfig.Process("", &Env); err != nil {
		log.Fatalf("[ERROR] Failed to process env: %s", err.Error())
	}

	durations := map[string]time.Duration{
		"1m": time.Minute,
		"1h": time.Hour,
		"1d": 24 * time.Hour,
	}

	Config.Durations = durations
	Config.LogFile = "gotrading.log"
	Config.ProductCode = "BTC_JPY"
	Config.TradeDuration = durations["1m"]
	Config.BackTest = true
	Config.UsePercent = 0.9
	Config.DataLimit = 365
	Config.StopLimitPercent = 0.8
	Config.NumRanking = 2
	Config.Deadline = 3600
	Config.SQLDriver = "mysql"
	Config.Port = 8090
}
