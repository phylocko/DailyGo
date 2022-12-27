/*
Copyright © 2022 Vladislav Pavkin [phylocko@gmail.com]
*/
package api

import (
	"dailygo/internal/ipinfo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InfoView(c *gin.Context) {

	target, err := ipinfo.GetInfo(c.ClientIP())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, target)
}
