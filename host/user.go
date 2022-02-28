package host

import (
	"hedgex-public/gl"
	"hedgex-public/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTraders(c *gin.Context) {
	contract := c.Query("contract")
	traders, _, err := model.GetUsers(contract)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"err_msg": "database error " + contract,
		})
		gl.OutLogger.Error("Get trader from database error : %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   traders,
	})
}

//GetKlineData get the contract's history kline data
func GetTradeRecords(c *gin.Context) {
	contract := c.Query("contract")
	account := c.Query("account")
	count, _ := strconv.Atoi(c.DefaultQuery("count", "1"))
	data, err := model.GetTradeRecords(contract, account, count)
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

func GetInterest(c *gin.Context) {
	contract := c.Query("contract")
	account := c.Query("account")
	count, _ := strconv.Atoi(c.DefaultQuery("count", "1"))
	interests, err := model.GetInterests(contract, account, count)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"err_msg": "database error " + contract,
		})
		gl.OutLogger.Error("Get interests from database error : %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   interests,
	})
}

func GetExplosive(c *gin.Context) {
	contract := c.Query("contract")
	account := c.Query("account")
	count, _ := strconv.Atoi(c.DefaultQuery("count", "1"))
	explosives, err := model.GetExplosive(contract, account, count)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"err_msg": "database error " + contract,
		})
		gl.OutLogger.Error("Get explosive from database error : %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   explosives,
	})
}
