package admin

import (
	"context"
	"dailygo/internal/database"
	"dailygo/internal/ipinfo"
	"dailygo/internal/models"
	"errors"
	"fmt"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexView(c *gin.Context) {

	context := gin.H{
		"title": "Main page",
	}
	c.HTML(http.StatusOK, "index.html", context)
}

func InfoView(c *gin.Context) {

	info, err := ipinfo.GetInfo(c.ClientIP())

	pageData := gin.H{}

	if err != nil {

		if errors.Is(err, context.DeadlineExceeded) {
			pageData["error"] = "ipinfo.io connection is timed out"
			c.HTML(http.StatusOK, "info.html", pageData)
			return

		} else {
			pageData["error"] = "Error communicating with ipinfo.io"
			c.HTML(http.StatusOK, "info.html", pageData)
			return

		}

	}

	pageData["info"] = info
	c.HTML(http.StatusOK, "info.html", pageData)

}

func ResourceView(c *gin.Context) {
	db := database.Db

	var resource models.Resource

	strID := c.Param("id")
	if strID == "" {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}

	intID, err := strconv.Atoi(strID)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}

	db.First(&resource, intID)

	if c.Request.Method == http.MethodPost {

		back := fmt.Sprintf("/resources/%d", intID)

		switch action := c.PostForm("action"); action {

		case "resolve_resource":

			resource.Resolve(db)
			c.Redirect(http.StatusMovedPermanently, back)
			return

		case "delete_address":

			strID := c.PostForm("address_id")

			intID, err := strconv.Atoi(strID)
			if err != nil {
				c.Redirect(http.StatusMovedPermanently, back)
				return
			}

			ip := models.IpAddress{}
			db.Where("resource_id = ?", resource.ID).First(&ip, intID)
			db.Delete(&ip)
			c.Redirect(http.StatusMovedPermanently, back)
			return

		}
	}

	var addresses []models.IpAddress

	db.Where("resource_id = ?", resource.ID).Find(&addresses)
	context := gin.H{"resource": resource, "addresses": addresses}
	c.HTML(http.StatusOK, "resource.html", context)
}

func ResourceListView(c *gin.Context) {
	db := database.Db

	var resources []models.Resource

	if err := db.Find(&resources).Error; err != nil {
		fmt.Println("Unable to get rows.", err)
	}

	context := gin.H{"resources": resources}

	if c.Request.Method == http.MethodPost {

		switch action := c.PostForm("action"); action {

		case "create_resource":
			domain := c.PostForm("domain")
			resource := models.Resource{Domain: domain}
			db.Create(&resource)
			resource.Resolve(db)
			c.Redirect(http.StatusMovedPermanently, "/resources/")
			return

		case "delete_resource":
			strID := c.PostForm("resource_id")

			intID, err := strconv.Atoi(strID)
			if err != nil {
				c.Redirect(http.StatusMovedPermanently, "/resources/")
				return
			}

			resource := models.Resource{}
			db.First(&resource, intID)

			db.Delete(&resource)
			c.Redirect(http.StatusMovedPermanently, "/resources/")
			return

		default:
			c.HTML(http.StatusNotFound, "404.html", nil)
			return
		}

	}

	c.HTML(http.StatusOK, "resource_list.html", context)
}
