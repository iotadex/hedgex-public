package host

import (
	"hedgex-public/config"
	"hedgex-public/gl"
	"hedgex-public/model"
	"hedgex-public/service"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/triplefi/go-logger/logger"
)

func StartHttpServer() {
	router := InitRouter()
	go router.Run(":" + strconv.Itoa(config.HttpPort))
}

// Index homepage
func PingPong(c *gin.Context) {
	ts, err := model.GetLatestContractUpdateTime()
	if err != nil {
		c.String(http.StatusOK, "mysql error")
		gl.OutLogger.Info("connect to mysql error. %v", err)
		return
	}
	if ts.Add(time.Hour).Before(time.Now()) {
		c.String(http.StatusOK, "event server error")
		gl.OutLogger.Info("event server error. %v", ts)
		return
	}
	res := "pong"
	for conAddr := range config.Contract {
		if atomic.LoadInt64(service.ChainNodeErr[conAddr]) != 0 {
			res = conAddr + " chain node error."
			break
		} else if atomic.LoadInt64(service.ContractPriceErr[conAddr]) != 0 {
			res = conAddr + " price was not update."
			break
		}
	}
	c.String(http.StatusOK, res)
}

// InitRouter init the router
func InitRouter() *gin.Engine {
	if err := os.MkdirAll("./logs/http", os.ModePerm); err != nil {
		log.Panic("Create dir './logs/http' error. " + err.Error())
	}
	GinLogger, err := logger.New("logs/http/gin.log", 2, 100*1024*1024, 10)
	if err != nil {
		log.Panic("Create GinLogger file error. " + err.Error())
	}

	router := gin.New()
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: GinLogger}), gin.Recovery())
	router.SetTrustedProxies(nil)

	router.GET("/api/ping", PingPong)

	//con
	con := router.Group("/api/contract")
	{
		con.GET("/trade_pairs", GetTradePairs)
		con.GET("/pair_params", GetPairParams)
		con.GET("/trade", GetTradeRecordsByContract)
		con.GET("/explosive", GetExplosiveRecordsByContract)
		con.GET("/kline", GetKlineData)
		con.GET("/position", GetStatPositions)
	}

	//acc
	acc := router.Group("/api/account")
	{
		acc.GET("/", GetTraders)
		acc.GET("/trade", GetTradeRecords)
		acc.GET("/interest", GetInterest)
		acc.GET("/explosive", GetExplosive)
	}

	//wss
	wss := router.Group("/wss")
	{
		wss.GET("/kline", klineSender)
	}

	other := router.Group("/api/odds")
	{
		other.GET("/add_email", AddEmail)
		other.GET("/emails", GetEmails)
		other.GET("/testcoin", SendTestCoins)
	}

	return router
}
