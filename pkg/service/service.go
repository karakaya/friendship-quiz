package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/karakaya/friendship-quiz/pkg/model"
	"github.com/karakaya/friendship-quiz/pkg/repository"
	"github.com/karakaya/friendship-quiz/pkg/request"
)

type QuizService interface {
	Create(ctx context.Context, req request.CreateQuizRequest) (interface{}, error)
	Get(ctx context.Context, id uuid.UUID) (interface{}, error)
}

type quizService struct {
	repo repository.Repository
}

func NewService(r repository.Repository) QuizService {
	return quizService{repo: r}
}

func (q quizService) Create(ctx context.Context, req request.CreateQuizRequest) (interface{}, error) {

	quiz := model.Quiz{
		ID:        uuid.New(),
		Questions: req.Questions,
	}

	return q.repo.Create(ctx, quiz)
}

func (q quizService) Get(ctx context.Context, id uuid.UUID) (interface{}, error) {
	return nil, nil
}
