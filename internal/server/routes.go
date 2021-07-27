package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(router *gin.RouterGroup) {
	router.GET("health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
		return
	})
}

func registerRoutes(router *gin.Engine) *gin.RouterGroup {
	v1 := router.Group("/api/v1/")
	{
		HealthCheck(v1)
		
		//TODO: add other routes
	}
	return v1
}
