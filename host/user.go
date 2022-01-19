package host

import (
	"encoding/json"
	"hedgex-public/gl"
	"hedgex-public/model"
	"net/http"
	"strconv"
)

//GetKlineData get the contract's history kline data
func GetTradeRecords(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			gl.OutLogger.Error("Panic: %v", err)
		}
	}()
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

func GetTraders(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			gl.OutLogger.Error("Panic: %v", err)
		}
	}()
	w.Header().Add("content-type", "application/json")
	contract := r.URL.Query().Get("contract")
	traders, _, err := model.GetUsers(contract)
	if err != nil {
		str, _ := json.Marshal(map[string]interface{}{
			"result":   false,
			"err_code": gl.DATABASE_ERROR,
			"err_msg":  "database error",
		})
		w.Write(str)
		gl.OutLogger.Error("Get trader from database error : %v", err)
		return
	}

	str, _ := json.Marshal(map[string]interface{}{
		"result": true,
		"data":   traders,
	})
	w.Write(str)
}

func GetInterest(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			gl.OutLogger.Error("Panic: %v", err)
		}
	}()
	w.Header().Add("content-type", "application/json")
	contract := r.URL.Query().Get("contract")
	account := r.URL.Query().Get("account")
	count, _ := strconv.Atoi(r.URL.Query().Get("count"))
	interests, err := model.GetInterests(contract, account, count)
	if err != nil {
		str, _ := json.Marshal(map[string]interface{}{
			"result":   false,
			"err_code": gl.DATABASE_ERROR,
			"err_msg":  "database error",
		})
		w.Write(str)
		gl.OutLogger.Error("Get interests from database error : %v", err)
		return
	}

	str, _ := json.Marshal(map[string]interface{}{
		"result": true,
		"data":   interests,
	})
	w.Write(str)
}

func GetExplosive(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			gl.OutLogger.Error("Panic: %v", err)
		}
	}()
	w.Header().Add("content-type", "application/json")
	contract := r.URL.Query().Get("contract")
	account := r.URL.Query().Get("account")
	count, _ := strconv.Atoi(r.URL.Query().Get("count"))
	interests, err := model.GetInterests(contract, account, count)
	if err != nil {
		str, _ := json.Marshal(map[string]interface{}{
			"result":   false,
			"err_code": gl.DATABASE_ERROR,
			"err_msg":  "database error",
		})
		w.Write(str)
		gl.OutLogger.Error("Get explosive from database error : %v", err)
		return
	}

	str, _ := json.Marshal(map[string]interface{}{
		"result": true,
		"data":   interests,
	})
	w.Write(str)
}

func SendTestCoins(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			gl.OutLogger.Error("Panic: %v", err)
		}
	}()
	w.Header().Add("content-type", "application/json")
	account := r.URL.Query().Get("user")
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

	go gl.SendTestCoins(account)

	str, _ := json.Marshal(map[string]interface{}{
		"result": true,
		"data":   "",
	})
	w.Write(str)
}
