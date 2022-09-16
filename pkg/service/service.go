package service

import (
	"context"

	"github.com/Olexander753/WB_L0/internal/schema"
	"github.com/Olexander753/WB_L0/pkg/repository"
)

type Model interface {
	InsertModel(ctx context.Context, model schema.Model) (string, error)
	SelectModel(ctx context.Context, uid string) (schema.Model, error)
}

type EventStore interface {
	Close()
	CreateModel() error
}

type Service struct {
	Model
}

func NewService(repo *repository.Storage) *Service {
	return &Service{
		Model: NewModelService(repo.Model),
	}
}
