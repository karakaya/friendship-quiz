package request

import "github.com/karakaya/friendship-quiz/pkg/model"

type CreateQuizRequest struct {
	Questions []model.Question `json:"question"`
}
