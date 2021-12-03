package host

import (
	"encoding/json"
	"hedgex-server/config"
	"hedgex-server/gl"
	"log"
	"net/http"
	"strconv"
)

func StartHttpServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", Index)
	mux.HandleFunc("/api/contract/kline", GetKlineData)
	mux.HandleFunc("/api/contract/trade_pairs", GetPairs)

	mux.HandleFunc("/api/account/trade", GetTradeRecords)
	mux.HandleFunc("/api/account/gettestcoin", SendTestCoins)

	mux.HandleFunc("/wss/kline", klineSender) // this is for websocket

	gl.HttpServer = &http.Server{
		Addr:    "localhost:" + strconv.Itoa(config.HttpPort),
		Handler: mux,
	}
	err := gl.HttpServer.ListenAndServe()
	if err != nil {
		log.Panic(err.Error())
	}
}

// Index homepage
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	str, _ := json.Marshal(map[string]interface{}{
		"result": true,
		"data":   "hedgex api",
	})
	w.Write(str)
}
