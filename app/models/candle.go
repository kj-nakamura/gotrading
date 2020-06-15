package models

import (
	"fmt"
	"gotrading/bitflyer"
	"gotrading/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Candle is for btc's candle
type Candle struct {
	ProductCode string        `json:"product_code"`
	Duration    time.Duration `json:"duration"`
	Time        time.Time     `json:"time"`
	Open        float64       `json:"open"`
	Close       float64       `json:"close"`
	High        float64       `json:"high"`
	Low         float64       `json:"low"`
	Volume      float64       `json:"volume"`
}

// NewCandle is 引数に準じたCandle情報を返す
func NewCandle(productCode string, duration time.Duration, timeDate time.Time, open, close, high, low, volume float64) *Candle {
	return &Candle{
		productCode,
		duration,
		timeDate,
		open,
		close,
		high,
		low,
		volume,
	}
}

func (c *Candle) TableName() string {
	return GetCandleTableName(c.ProductCode, c.Duration)
}

func (c *Candle) Create() {
	cmd := fmt.Sprintf("INSERT INTO %s (time, open, close, high, low, volume) VALUES (?, ?, ?, ?, ?, ?)", c.TableName())
	DbConnection.Exec(cmd, c.Time, c.Open, c.Close, c.High, c.Low, c.Volume)
}

func (c *Candle) Save() {
	DbConnection.Table(c.TableName()).Where("time = ?", c.Time).Update(map[string]interface{}{
		"Open":   c.Open,
		"Close":  c.Close,
		"High":   c.High,
		"Low":    c.Low,
		"Volume": c.Volume,
	})
}

// GetCandle is 引数のTimeに合致したキャンドルを返す
func GetCandle(productCode string, duration time.Duration, dateTime time.Time) *Candle {
	tableName := GetCandleTableName(productCode, duration)
	row := DbConnection.Table(tableName).Select([]string{"Time", "Open", "Close", "High", "Low", "Volume"}).Where("time = ?", dateTime).Row()

	var candle Candle
	// Candle structに取得したクエリを追加する
	err := row.Scan(&candle.Time, &candle.Open, &candle.Close, &candle.High, &candle.Low, &candle.Volume)
	if err != nil {
		return nil
	}
	return NewCandle(productCode, duration, candle.Time, candle.Open, candle.Close, candle.High, candle.Low, candle.Volume)
}

// CreateCandleWithDuration is Durationごとにキャンドルを更新する
func CreateCandleWithDuration(ticker bitflyer.Ticker, productCode string, duration time.Duration) bool {
	currentCandle := GetCandle(productCode, duration, ticker.TruncateDateTime(duration))
	price := ticker.GetMidPrice()
	// Durationごとの実行時はキャンドルを作成する
	if currentCandle == nil {
		candle := NewCandle(productCode, duration, ticker.TruncateDateTime(duration),
			price, price, price, price, ticker.Volume)
		candle.Create()
		return true
	}

	if currentCandle.High <= price {
		currentCandle.High = price
	} else if currentCandle.Low >= price {
		currentCandle.Low = price
	}
	currentCandle.Volume += ticker.Volume
	currentCandle.Close = price
	currentCandle.Save()

	t := time.Now()
	t = t.Add(time.Duration(-config.Config.Deadline) * time.Minute)

	DeleteCandleWithDuration(productCode, duration, t)

	return false
}

func DeleteCandleWithDuration(productCode string, duration time.Duration, deadline time.Time) bool {
	tableName := GetCandleTableName(productCode, duration)
	DbConnection.Table(tableName).Unscoped().Where("time < ?", deadline).Delete(Duration{})

	return true
}

func GetAllCandle(productCode string, duration time.Duration, limit int) (dfCandle *DataFrameCandle, err error) {
	tableName := GetCandleTableName(productCode, duration)

	// descでlimit数レコードを取得し、ascに並べ替える
	cmd := DbConnection.Table(tableName).Select([]string{"time", "open", "close", "high", "low", "volume"}).Order("time desc").Limit(limit).SubQuery()

	rows, err := DbConnection.Raw("SELECT * FROM ? AS cmd ORDER BY time ASC;", cmd).Rows()

	if err != nil {
		return
	}

	defer rows.Close()

	dfCandle = &DataFrameCandle{}
	dfCandle.ProductCode = productCode
	dfCandle.Duration = duration
	for rows.Next() {
		var candle Candle
		candle.ProductCode = productCode
		candle.Duration = duration
		rows.Scan(&candle.Time, &candle.Open, &candle.Close, &candle.High, &candle.Low, &candle.Volume)
		dfCandle.Candles = append(dfCandle.Candles, candle)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return dfCandle, nil
}
