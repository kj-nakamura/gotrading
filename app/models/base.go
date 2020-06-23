package models

import (
	"fmt"

	"gotrading/config"
	"log"
	"time"

	// orm
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	tableNameSignalEvents = "signal_events"
)

// DbConnection is for using global
var DbConnection *gorm.DB

// GetCandleTableName is print productCode and duration
func GetCandleTableName(productCode string, duration time.Duration) string {
	return fmt.Sprintf("%s_%s", productCode, duration)
}

// Event is table of events for buy or sell
type Event struct {
	gorm.Model
	Time        time.Time `gorm:"primary_key: not null"`
	ProductCode string
	Side        string
	Price       float64
	Size        float64
}

// Duration is table of events for buy or sell
type Duration struct {
	gorm.Model
	Time   *time.Time `gorm:"primary_key: not null"`
	Open   float64
	Close  float64
	High   float64
	Low    float64
	Volume float64
}

func init() {
	var err error
	DbConnection, err = gorm.Open(config.Config.SQLDriver, config.Config.DbUserName+":"+config.Config.DbPassword+"@tcp(mysql:3306)/"+config.Config.DbName+"?charset=utf8&parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatalln(err)
	}

	DbConnection.Table(tableNameSignalEvents).AutoMigrate(&Event{})
	for _, duration := range config.Config.Durations {
		tableName := GetCandleTableName(config.Config.ProductCode, duration)
		DbConnection.Table(tableName).AutoMigrate(&Duration{})
	}
}
