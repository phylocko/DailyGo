/*
Copyright © 2022 Vladislav Pavkin [phylocko@gmail.com]
*/
package api

import (
	"dailygo/internal/ipinfo"
	proto_ipinfo "dailygo/proto/ipinfo"
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

func InfoPBView(c *gin.Context) {

	target, err := ipinfo.GetInfo(c.ClientIP())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	pResonse := &proto_ipinfo.IPInfoData{
		Ip: target.Ip,

		City:     target.City,
		Region:   target.Region,
		Country:  target.Country,
		Loc:      target.Loc,
		Org:      target.Org,
		Postal:   target.Postal,
		Timezone: target.Timezone,
		Bogon:    target.Bogon,
	}

	c.ProtoBuf(http.StatusOK, pResonse)
}
