package contracts

type GetQuestionResponse struct {
	Questions []string `json:"questions" `
}

type AddQuestionRequest struct {
	Link string `json:"link" binding:"required"`
}

type AddQuestionResponse struct {
	Message string `json:"message"`
}

type RemoveQuestionRequest struct {
	Link string `json:"link" binding:"required"`
}

type RemoveQuestionResponse struct {
	Message string `json:"message"`
}

type ChangeStatusOfQuestionRequest struct {
	QuestionLink string `json:"link"  binding:"required"`
	Status       bool   `json:"status" binding:"required"`
}

type ChangeStatusOfQuestionResponse struct {
	Message string `json:"message"`
}
