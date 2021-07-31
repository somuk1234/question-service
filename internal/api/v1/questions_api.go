package question

import (
	"log"
	"net/http"
	"questions-keeper-service/internal/contracts"
	"questions-keeper-service/internal/drivers"

	"github.com/gin-gonic/gin"
)

type GetToDoQuestionsApi struct {
	Resp *contracts.GetQuestionResponse
}

func GetToDoQuestions(router *gin.RouterGroup) {
	router.GET("/todo-questions/", func(ctx *gin.Context) {
		api := &GetToDoQuestionsApi{
			Resp: new(contracts.GetQuestionResponse),
		}
		var err error
		api.Resp.Questions, err = drivers.GetToDoQuestions()
		//TODO: add response writter in a better way
		var res = make(map[string]interface{})
		res["data"] = api.Resp.Questions
		res["message"] = "Request Succesful"
		if err != nil {
			res["message"] = "Request Failed"
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		ctx.JSON(http.StatusOK, res)
	})
}

type GetDoneQuestionsApi struct {
	Resp *contracts.GetQuestionResponse
}

func GetDoneQuestions(router *gin.RouterGroup) {
	router.GET("/done-questions/", func(ctx *gin.Context) {
		api := &GetDoneQuestionsApi{
			Resp: new(contracts.GetQuestionResponse),
		}
		var err error
		api.Resp.Questions, err = drivers.GetDoneQuestions()
		//TODO: add response writter in a better way
		var res = make(map[string]interface{})
		res["data"] = api.Resp.Questions
		res["message"] = "Request Succesful"
		if err != nil {
			res["message"] = "Request Failed"
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		ctx.JSON(http.StatusOK, res)
	})
}

type AddQuestionApi struct {
	Req  *contracts.AddQuestionRequest
	Resp *contracts.AddQuestionResponse
}

func AddQuestion(router *gin.RouterGroup) {
	router.POST("/add-question/", func(ctx *gin.Context) {
		api := &AddQuestionApi{
			Req:  new(contracts.AddQuestionRequest),
			Resp: new(contracts.AddQuestionResponse),
		}
		err := api.handleValidation(ctx)
		var res = make(map[string]interface{})
		if err != nil {
			res["message"] = "Invalid Request"
			res["data"] = ""
			ctx.JSON(http.StatusBadRequest, res)
			return
		}
		err = drivers.AddQuestion(api.Req.Link)
		if err != nil {
			res["message"] = err.Error()
			res["data"] = ""
			ctx.JSON(http.StatusBadRequest, res)
			return
		}
		res["message"] = "Added succesfully"
		res["data"] = ""
		ctx.JSON(http.StatusOK, res)

	})
}

type RemoveQuestionApi struct {
	Req  *contracts.RemoveQuestionRequest
	Resp *contracts.RemoveQuestionResponse
}

func RemoveQuestion(router *gin.RouterGroup) {
	router.POST("/remove-question/", func(ctx *gin.Context) {
		api := &RemoveQuestionApi{
			Req:  new(contracts.RemoveQuestionRequest),
			Resp: new(contracts.RemoveQuestionResponse),
		}
		err := api.handleValidation(ctx)
		var res = make(map[string]interface{})
		if err != nil {
			res["message"] = "Invalid Request"
			res["data"] = ""
			ctx.JSON(http.StatusBadRequest, res)
			return
		}
		err = drivers.RemoveQuestion(api.Req.Link)
		if err != nil {
			res["message"] = err.Error()
			res["data"] = ""
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, res)
			return
		}
		res["message"] = "Removed succesfully"
		res["data"] = ""
		ctx.JSON(http.StatusOK, res)
	})
}

type ChangeStatusOfQuestionApi struct {
	Req  *contracts.ChangeStatusOfQuestionRequest
	Resp *contracts.ChangeStatusOfQuestionResponse
}

func ChangeStatusOfQuestion(router *gin.RouterGroup) {
	router.PUT("/change-status/", func(ctx *gin.Context) {
		api := &ChangeStatusOfQuestionApi{
			Req:  new(contracts.ChangeStatusOfQuestionRequest),
			Resp: new(contracts.ChangeStatusOfQuestionResponse),
		}
		err := api.handleValidation(ctx)
		var res = make(map[string]interface{})
		if err != nil {
			res["message"] = "Invalid Request"
			res["data"] = ""
			ctx.JSON(http.StatusBadRequest, res)
			return
		}
		err = drivers.ChangeStatusOfQuestion(api.Req.QuestionLink, api.Req.Status)
		if err != nil {
			res["message"] = err.Error()
			res["data"] = ""
			ctx.JSON(http.StatusBadRequest, res)
			return
		}
		res["message"] = "Status Chnaged succesfully"
		res["data"] = ""
		ctx.JSON(http.StatusOK, res)
	})
}

func (api *AddQuestionApi) handleValidation(ctx *gin.Context) error {
	err := ctx.ShouldBindJSON(&api.Req)
	log.Println(api.Req)
	if err != nil {
		return err
	}
	return nil
}

func (api *RemoveQuestionApi) handleValidation(ctx *gin.Context) error {
	err := ctx.ShouldBindJSON(&api.Req)
	if err != nil {
		return err
	}
	return nil
}

func (api *ChangeStatusOfQuestionApi) handleValidation(ctx *gin.Context) error {
	err := ctx.ShouldBindJSON(&api.Req)
	if err != nil {
		return err
	}
	return nil
}
