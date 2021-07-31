package repository

import (
	managers "questions-keeper-service/internal/manager"
	"questions-keeper-service/internal/models"

	"github.com/gofrs/uuid"
)

func GetQuestionByQuestionLink(Qlink string) (models.Question, error) {
	var Question models.Question
	result := managers.Database.Where("link = ?", Qlink).First(&Question)
	if result.Error != nil {
		return Question, result.Error
	}
	return Question, nil
}

func DeleteQuestion(Qlink string) error {
	var Question models.Question
	result := managers.Database.Where("link = ?", Qlink).Delete(&Question)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetToDoQuestions() ([]models.Question, error) {
	var Questions []models.Question
	result := managers.Database.Where("status = ?", false).Find(&Questions)
	if result.Error != nil {
		return Questions, result.Error
	}
	return Questions, nil
}

func GetDoneQuestions() ([]models.Question, error) {
	var Questions []models.Question
	result := managers.Database.Where("status = ?", true).Find(&Questions)
	if result.Error != nil {
		return Questions, result.Error
	}
	return Questions, nil
}

func ChangeStatusOfQuestion(question models.Question) error {
	result := managers.Database.Save(&question)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AddQuestion(Qlink string) error {
	var Question models.Question
	Question.ID, _ = uuid.NewV4()
	Question.Status = false
	Question.Rating = 1000
	Question.Link = Qlink
	result := managers.Database.Create(&Question)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
