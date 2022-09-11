package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/karakaya/friendship-quiz/pkg/repository"
)

type QuizService interface {
	Create(ctx context.Context, obj interface{}) error
	Get(ctx context.Context, id uuid.UUID) (interface{}, error)
}

type quizService struct {
	repo repository.Repository
}

func NewService(r repository.Repository) QuizService {
	return quizService{repo: r}
}

func (q quizService) Create(ctx context.Context, obj interface{}) error {
	return q.repo.Create(ctx, obj)

}

func (q quizService) Get(ctx context.Context, id uuid.UUID) (interface{}, error) {
	return nil, nil
}
