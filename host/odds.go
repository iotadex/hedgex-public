package host

import (
	"hedgex-public/config"
	"hedgex-public/gl"
	"hedgex-public/model"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Index homepage
func AddEmail(c *gin.Context) {
	email := c.Query("email")
	if !isEmailValid(email) {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"err_msg": "invalid email" + email,
		})
		gl.OutLogger.Warn("Invalid email : %s", email)
		return
	}

	ip := c.ClientIP()

	if count, err := model.GetIpCount(ip); err != nil || count > config.IpLimit {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"err_msg": "send email fail",
		})
		gl.OutLogger.Warn("Get ip count error. %s : %d : %s : %v", ip, count, email, err)
		return
	}

	if err := model.InsertEmail(email, ip); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"err_msg": "invalid email",
		})
		gl.OutLogger.Warn("Insert email error. %s : %v", email, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   "",
	})
}

func GetEmails(c *gin.Context) {
	from, _ := strconv.ParseInt(c.DefaultQuery("from", "0"), 10, 64)
	to, _ := strconv.ParseInt(c.DefaultQuery("to", "0"), 10, 64)
	if to == 0 {
		to = time.Now().Unix()
	}
	emails, err := model.GetEmails(time.Unix(from, 0).Format("2006-01-02 15:04:05"), time.Unix(to, 0).Format("2006-01-02 15:04:05"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"err_msg": "db error",
		})
		gl.OutLogger.Error("Get emails from db error. %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   emails,
	})
}

// isEmailValid checks if the email provided passes the required structure
// and length test. It also checks the domain has a valid MX record.
func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	if !emailRegex.MatchString(e) {
		return false
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}
