package main

import (
	"dailygo/internal/web/admin"
	"dailygo/internal/web/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	apiRoutes := r.Group("/api")
	{
		apiRoutes.GET("/info", api.InfoView)
	}

	adminRoutes := r.Group("/")
	{
		adminRoutes.GET("/", admin.IndexView)
		adminRoutes.GET("/info", admin.InfoView)
	}

	r.Run()
}
