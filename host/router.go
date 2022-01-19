package host

import (
	"encoding/json"
	"hedgex-public/config"
	"hedgex-public/gl"
	"log"
	"net/http"
	"strconv"
)

func StartHttpServer() {
	mux := http.NewServeMux()
	chainNet := ""
	if len(config.ChainNode.Name) > 0 {
		chainNet = "/" + config.ChainNode.Name
	}
	mux.HandleFunc(chainNet+"/api", Index)
	mux.HandleFunc(chainNet+"/api/contract/kline", GetKlineData)
	mux.HandleFunc(chainNet+"/api/contract/trade_pairs", GetPairs)
	mux.HandleFunc(chainNet+"/api/contract/position", GetStatPositions)

	mux.HandleFunc(chainNet+"/api/account", GetTraders)
	mux.HandleFunc(chainNet+"/api/account/trade", GetTradeRecords)
	mux.HandleFunc(chainNet+"/api/account/interest", GetInterest)
	mux.HandleFunc(chainNet+"/api/account/explosive", GetExplosive)
	mux.HandleFunc(chainNet+"/api/account/gettestcoin", SendTestCoins)

	mux.HandleFunc(chainNet+"/wss/kline", klineSender) // this is for websocket

	gl.HttpServer = &http.Server{
		Addr:    "localhost:" + strconv.Itoa(config.HttpPort),
		Handler: mux,
	}
	err := gl.HttpServer.ListenAndServe()
	if err != nil {
		log.Println(err.Error())
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
