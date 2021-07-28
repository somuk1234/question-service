package question

import "github.com/gin-gonic/gin"

func GetToDoQuestions(router *gin.RouterGroup) {
	router.GET("/todo-questions/", func(ctx *gin.Context) {
		//TODO: write login here
	})
}

func GetDoneQuestions(router *gin.RouterGroup) {
	router.GET("/done-questions/", func(ctx *gin.Context) {
		//TODO: write login here
	})
}

func AddQuestion(router *gin.RouterGroup) {
	router.POST("/add-question/", func(ctx *gin.Context) {
		//TODO: write logic here
	})
}
func RemoveQuestion(router *gin.RouterGroup) {
	router.POST("/remove-question/", func(ctx *gin.Context) {
		//TODO: write logic here
	})
}

func UpdateQuestion(router *gin.RouterGroup) {
	router.PUT("/update-question/", func(ctx *gin.Context) {
		//TODO: write logic here
	})
}

func ChangeStatusOfQuestion(router *gin.RouterGroup) {
	router.PUT("/change-status/", func(ctx *gin.Context) {
		//TODO: write logic here
	})
}
