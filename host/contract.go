package host

import (
	"encoding/json"
	"hedgex-server/config"
	"hedgex-server/gl"
	"hedgex-server/model"
	"net/http"
	"strconv"
)

type pair struct {
	Contract     string `json:"contract"`
	MarginCoin   string `json:"margin_coin"`
	TradeCoin    string `json:"trade_coin"`
	DayOpenPrice int64  `json:"open_price"`
	IndexPrice   int64  `json:"index_price"`
}

//GetPairs get the contract's trade pairs
func GetPairs(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	//get current indexPrice and current day's open price
	pairs := make([]pair, len(config.Contract))
	for i := range config.Contract {
		pairs[i].Contract = config.Contract[i].Address
		pairs[i].MarginCoin = config.Contract[i].MarginCoin
		pairs[i].TradeCoin = config.Contract[i].TradeCoin
		if skd := gl.CurrentKlineDatas[config.Contract[i].Address]; skd != nil {
			candle := skd.GetCurrent("d1")
			pairs[i].DayOpenPrice = candle[0]
			pairs[i].IndexPrice = candle[3]
		}
	}

	str, _ := json.Marshal(map[string]interface{}{
		"result": true,
		"data":   pairs,
	})
	w.Write(str)
}

//GetKlineData get the contract's history kline data
func GetKlineData(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	contract := r.URL.Query().Get("contract")
	t := r.URL.Query().Get("type")
	count, _ := strconv.Atoi(r.URL.Query().Get("count"))
	if _, exist := gl.CurrentKlineDatas[contract]; !exist {
		str, _ := json.Marshal(map[string]interface{}{
			"result": false,
			"data":   "Contract not exist. " + contract,
		})
		gl.OutLogger.Warn("Contract not exist. " + contract)
		w.Write(str)
		return
	}
	data := gl.CurrentKlineDatas[contract].Get(t, count)
	str, _ := json.Marshal(map[string]interface{}{
		"result": true,
		"data":   data,
	})
	w.Write(str)
}

func GetStatPositions(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	data, err := model.GetStatPositions()
	if err != nil {
		str, _ := json.Marshal(map[string]interface{}{
			"result": false,
			"data":   "Get Position Stat From DB Error. " + err.Error(),
		})
		w.Write(str)
		gl.OutLogger.Warn("Get Position Stat From DB Error. %v", err)
		return
	}

	str, _ := json.Marshal(map[string]interface{}{
		"result": true,
		"data":   data,
	})
	w.Write(str)
}
