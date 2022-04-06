package host

import (
	"hedgex-public/config"
	"hedgex-public/gl"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadShareResultImg(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": false,
			"err":    "get file error",
		})
		gl.OutLogger.Error("Upload file error. %v", err)
		return
	}
	ext := path.Ext(file.Filename)
	fileName := strconv.FormatInt(time.Now().UnixNano(), 10) + ext
	filePath := filepath.Join(config.Upload, "/", fileName)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": false,
			"err":    "save file error",
		})
		gl.OutLogger.Error("Save file error. %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   fileName,
	})
}
