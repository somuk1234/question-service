package server

import (
	"net/http"
	question "questions-keeper-service/internal/api/v1"

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
		question.GetDoneQuestions(v1)
		question.AddQuestion(v1)
		question.ChangeStatusOfQuestion(v1)
		question.GetToDoQuestions(v1)
		question.RemoveQuestion(v1)
		question.UpdateQuestion(v1)
	}
	return v1
}
