package host

import (
	"encoding/json"
	"hedgex-server/config"
	"hedgex-server/gl"
	"hedgex-server/model"
	"net/http"
	"strconv"
)

//GetPairs get the contract's trade pairs
func GetPairs(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	pairs, err := model.GetPairs()
	if err != nil {
		str, _ := json.Marshal(map[string]interface{}{
			"result":   false,
			"err_code": gl.DATABASE_ERROR,
			"err_msg":  "database error : ",
		})
		w.Write(str)
		gl.OutLogger.Error("database error : %v", err)
		return
	}

	//get current indexPrice and current day's open price
	for i := range pairs {
		candle := gl.CurrentKlineDatas[pairs[i].Contract].GetCurrent("d1")
		pairs[i].DayOpenPrice = candle[0]
		pairs[i].IndexPrice = candle[3]
	}

	//update the config.Contract.Pairs. get current indexPrice and current day's open price
	if len(pairs) > 0 {
		contracts := make([]string, len(pairs))
		for i := range pairs {
			candle := gl.CurrentKlineDatas[pairs[i].Contract].GetCurrent("d1")
			pairs[i].DayOpenPrice = candle[0]
			pairs[i].IndexPrice = candle[3]
			contracts[i] = pairs[i].Contract
		}
		config.Contract.Pair = contracts
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
			"result": true,
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
