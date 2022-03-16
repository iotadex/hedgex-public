package kline

import "hedgex-public/config"

var (
	KlineTypes     []string
	KlineTimeCount map[string]int64
	DefaultDrivers map[string]KlineDriver
)

func init() {
	KlineTypes = []string{"m1", "m5", "m10", "m15", "m30", "h1", "h2", "h4", "h6", "h12", "d1"}
	KlineTimeCount = make(map[string]int64)
	KlineTimeCount["m1"] = 60
	KlineTimeCount["m5"] = 300
	KlineTimeCount["m10"] = 600
	KlineTimeCount["m15"] = 900
	KlineTimeCount["m30"] = 1800
	KlineTimeCount["h1"] = 3600
	KlineTimeCount["h2"] = 7200
	KlineTimeCount["h4"] = 14400
	KlineTimeCount["h6"] = 21600
	KlineTimeCount["h12"] = 43200
	KlineTimeCount["d1"] = 86400

	DefaultDrivers = make(map[string]KlineDriver)
	for conAddr := range config.Contract {
		if len(config.Redis.Addr) > 0 {
			DefaultDrivers[conAddr] = NewRedisKline(conAddr[2:6])
		} else {
			DefaultDrivers[conAddr] = NewMemoryKline(conAddr[2:6])
		}
	}
}

type KlineDriver interface {
	Get(kt string, count int) ([][5]int64, error)
	GetCurrent(kt string) ([5]int64, error)
	Append(kt string, currentData [5]int64) error
}
