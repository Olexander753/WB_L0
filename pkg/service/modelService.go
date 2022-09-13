package service

import (
	"context"

	"github.com/Olexander753/WB_L0/internal/schema"
	"github.com/Olexander753/WB_L0/pkg/repository"
)

type ModelService struct {
	repo repository.Model
}

func NewModelService(repo repository.Model) *ModelService {
	return &ModelService{repo: repo}
}

func (s *ModelService) InsertModel(ctx context.Context, model schema.Model) (string, error) {
	return s.repo.InsertModel(ctx, model)
}

func (s *ModelService) SelectModel(ctx context.Context, uid string) (schema.Model, error) {
	return s.repo.SelectModel(ctx, uid)
}
