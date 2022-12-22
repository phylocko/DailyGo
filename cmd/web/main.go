package main

import (
	"dailygo/internal/database"
	"dailygo/internal/models"
	"dailygo/internal/web/admin"
	"dailygo/internal/web/api"

	"github.com/gin-gonic/gin"
)

func runServer() {
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

		adminRoutes.GET("/resources", admin.ResourceListView)
		adminRoutes.POST("/resources", admin.ResourceListView)

		adminRoutes.GET("/resources/:id", admin.ResourceView)
		adminRoutes.POST("/resources/:id", admin.ResourceView)
	}

	r.Run()
}

// Migrate creates a database structure.
// Call it manually if needed.
func Migrate() {
	database.Db.AutoMigrate(
		&models.Resource{},
		&models.IpAddress{},
	)
}

func main() {
	database.CreateConnection()
	runServer()

}
