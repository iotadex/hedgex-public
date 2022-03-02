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
	i := 0
	for addr := range config.Contract {
		pairs[i].Contract = addr
		pairs[i].MarginCoin = config.Contract[addr].MarginCoin
		pairs[i].TradeCoin = config.Contract[addr].TradeCoin
		if skd := gl.CurrentKlineDatas[addr]; skd != nil {
			candle := skd.GetCurrent("d1")
			pairs[i].DayOpenPrice = candle[0]
			pairs[i].IndexPrice = candle[3]
		}
		i++
	}

	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   pairs,
	})
}

//GetTradePairs get the contract's trade pairs
func GetPairParams(c *gin.Context) {
	contract := c.Query("contract")
	//get current indexPrice and current day's open price
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   config.Contract[contract].Params,
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

func GetTradeRecordsByContract(c *gin.Context) {
	contract := c.Query("contract")
	count, _ := strconv.Atoi(c.DefaultQuery("count", "1"))
	if count > 200 {
		count = 200
	}
	data, err := model.GetTradeRecordsByContract(contract, count)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"err_msg": "database error " + contract,
		})
		gl.OutLogger.Error("Get trade records from database error : %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   data,
	})
}
