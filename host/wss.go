package host

import (
	"hedgex-public/config"
	"hedgex-public/gl"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func klineSender(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.Write([]byte("http upgrade to ws error :" + err.Error()))
		gl.OutLogger.Error("http upgrade to ws error :" + err.Error())
		return
	}
	defer c.Close()

	mt, message, err := c.ReadMessage()
	if err != nil || mt != websocket.TextMessage {
		c.WriteMessage(mt, message)
		gl.OutLogger.Error("Read msg error : %d, %s", mt, err.Error())
		return
	}

	strs := strings.Split(string(message), ":")
	gl.OutLogger.Info("Start wss client connection. %v", strs)
	ticker := time.NewTicker(time.Second * config.WsTick)
	defer ticker.Stop()
	for range ticker.C {
		if ckd, exist := gl.CurrentKlineDatas[strs[0]]; !exist {
			c.WriteJSON("contract no exist")
			break
		} else {
			data := ckd.GetCurrent(strs[1])
			if err := c.WriteJSON(data); err != nil {
				gl.OutLogger.Info("Write to ws error. %v", err)
				break
			}
		}
	}
}
