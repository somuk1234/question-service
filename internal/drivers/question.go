package drivers

import (
	"errors"
	"questions-keeper-service/internal/repository"
)

func GetToDoQuestions() ([]string, error) {
	Questions, err := repository.GetToDoQuestions()
	if err != nil {
		return []string{}, err
	}
	var result []string
	for _, val := range Questions {
		result = append(result, val.Link)
	}
	return result, nil
}

func GetDoneQuestions() ([]string, error) {
	Questions, err := repository.GetDoneQuestions()
	if err != nil {
		return []string{}, err
	}
	var result []string
	for _, val := range Questions {
		result = append(result, val.Link)
	}
	return result, nil
}

func AddQuestion(Qlink string) error {
	Question, _ := repository.GetQuestionByQuestionLink(Qlink)
	if Question.Link != "" {
		return errors.New("already exist")
	}
	return repository.AddQuestion(Qlink)
}
func RemoveQuestion(Qlink string) error {
	Question, _ := repository.GetQuestionByQuestionLink(Qlink)
	if Question.Link == "" {
		return errors.New("not exist")
	}
	return repository.DeleteQuestion(Qlink)
}

func ChangeStatusOfQuestion(Qlink string, status bool) error {
	Question, _ := repository.GetQuestionByQuestionLink(Qlink)
	if Question.Link == "" {
		return errors.New("not exist")
	}
	Question.Status = status
	return repository.ChangeStatusOfQuestion(Question)
}
