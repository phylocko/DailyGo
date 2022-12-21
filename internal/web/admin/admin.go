package admin

import (
	"dailygo/internal/ipinfo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexView(c *gin.Context) {

	context := gin.H{
		"title": "Main page",
	}
	c.HTML(http.StatusOK, "index.html", context)
}

func InfoView(c *gin.Context) {

	context, err := ipinfo.GetInfo(c.ClientIP())
	if err != nil {
		c.HTML(http.StatusOK, "info.html", gin.H{"error": err})
		return
	}

	c.HTML(http.StatusOK, "info.html", context)

}
