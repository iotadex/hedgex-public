package host

import (
	"hedgex-public/config"
	"hedgex-public/gl"
	"hedgex-public/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type pair struct {
	Contract     string `json:"contract"`
	MarginCoin   string `json:"margin_coin"`
	TradeCoin    string `json:"trade_coin"`
	DayOpenPrice int64  `json:"open_price"`
	IndexPrice   int64  `json:"index_price"`
}

//GetTradePairs get the contract's trade pairs
func GetTradePairs(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   pairs,
	})
}

//GetKlineData get the contract's history kline data
func GetKlineData(c *gin.Context) {
	contract := c.Query("contract")
	t := c.Query("type")
	count, _ := strconv.Atoi(c.DefaultQuery("count", "1"))
	if _, exist := gl.CurrentKlineDatas[contract]; !exist {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"err_msg": "Contract not exist. " + contract,
		})
		gl.OutLogger.Warn("Contract not exist. " + contract)
		return
	}
	data := gl.CurrentKlineDatas[contract].Get(t, count)
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   data,
	})
}

func GetStatPositions(c *gin.Context) {
	data, err := model.GetStatPositions()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"err_msg": "Get Position Stat From DB Error. " + err.Error(),
		})
		gl.OutLogger.Warn("Get Position Stat From DB Error. %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   data,
	})
}
