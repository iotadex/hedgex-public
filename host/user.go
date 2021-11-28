package host

import (
	"encoding/json"
	"hedgex-server/gl"
	"hedgex-server/model"
	"hedgex-server/service"
	"net/http"
	"strconv"
)

//GetKlineData get the contract's history kline data
func GetTradeRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	contract := r.URL.Query().Get("contract")
	account := r.URL.Query().Get("account")
	count, _ := strconv.Atoi(r.URL.Query().Get("count"))
	data, err := model.GetTradeRecords(contract, account, count)
	if err != nil {
		str, _ := json.Marshal(map[string]interface{}{
			"result":   false,
			"err_code": gl.DATABASE_ERROR,
			"err_msg":  "database error",
		})
		w.Write(str)
		gl.OutLogger.Error("Get trade records from database error : %v", err)
		return
	}

	str, _ := json.Marshal(map[string]interface{}{
		"result": true,
		"data":   data,
	})
	w.Write(str)
}

func SendTestCoins(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	account := r.URL.Query().Get("account")
	if err := model.UpdateTestCoin(account); err != nil {
		str, _ := json.Marshal(map[string]interface{}{
			"result":   false,
			"err_code": gl.DATABASE_ERROR,
			"err_msg":  "over count",
		})
		w.Write(str)
		gl.OutLogger.Error("Get trade records from database error : %v", err)
		return
	}

	go service.SendTestCoins(account)

	str, _ := json.Marshal(map[string]interface{}{
		"result": true,
		"data":   "",
	})
	w.Write(str)
}
