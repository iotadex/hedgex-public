package host

import (
	"hedgex-public/config"
	"hedgex-public/gl"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func klineSender(c *gin.Context) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"err_msg": "http upgrade to ws error : " + err.Error(),
		})
		gl.OutLogger.Error("http upgrade to ws error :" + err.Error())
		return
	}

	defer ws.Close()
	mt, message, err := ws.ReadMessage()
	if err != nil || mt != websocket.TextMessage {
		ws.WriteMessage(mt, message)
		gl.OutLogger.Error("Read msg error : %d, %s", mt, err.Error())
		return
	}

	strs := strings.Split(string(message), ":")
	gl.OutLogger.Info("Start wss client connection. %v", strs)
	ticker := time.NewTicker(time.Second * config.WsTick)
	defer ticker.Stop()
	for range ticker.C {
		if ckd, exist := gl.CurrentKlineDatas[strs[0]]; !exist {
			ws.WriteJSON("contract no exist")
			break
		} else {
			data := ckd.GetCurrent(strs[1])
			if err := ws.WriteJSON(data); err != nil {
				gl.OutLogger.Info("Write to ws error. %v", err)
				break
			}
		}
	}
}
