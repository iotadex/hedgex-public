package model

import (
	"hedgex-server/gl"
	"strconv"
)

func GetKlineData(contract string, klineType string, count int) ([][5]int64, error) {
	rows, err := gl.DB.Query("SELECT id,open,high,low,close FROM kline where contract='" + contract + "' and kline_type='" + klineType + "' order by id desc limit " + strconv.Itoa(count))
	if err != nil {
		return nil, err
	}
	candles := make([][5]int64, 0, count)
	for rows.Next() {
		var candle [5]int64
		rows.Scan(&candle[4], &candle[0], &candle[1], &candle[2], &candle[3])
		candles = append(candles, candle)
	}
	return candles, nil
}

func ReplaceKlineData(contract string, klineType string, candle [5]int64) error {
	_, err := gl.DB.Exec("replace into kline(contract,kline_type,id,open,high,low,close) values(?,?,?,?,?,?,?)", contract, klineType, candle[4], candle[0], candle[1], candle[2], candle[3])
	return err
}
