package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	ApiKey           string        `required:"true" split_words:"true"`
	ApiSecret        string        `required:"true" split_words:"true"`
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
	DbName           string `required:"true" split_words:"true"`
	DbUserName       string `required:"true" split_words:"true"`
	DbPassword       string `required:"true" split_words:"true"`
	SQLDriver        string `required:"true" split_words:"true"`
	Port             int    `required:"true" split_words:"true"`
}

var Config EnvConfig

func init() {
	if err := envconfig.Process("", &Config); err != nil {
		log.Fatalf("[ERROR] Failed to process env: %s", err.Error())
	}

	durations := map[string]time.Duration{
		"1m": time.Minute,
		"1h": time.Hour,
		"1d": 24 * time.Hour,
	}

	Config.Durations = durations
	Config.TradeDuration = durations[Config.TradeDuration.String()]
}
