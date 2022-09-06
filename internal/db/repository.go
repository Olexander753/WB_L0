package db

import (
	"context"

	"github.com/Olexander753/WB_L0/internal/schema"
)

type Repository interface {
	Close()
	InsertModel(ctx context.Context, model schema.Model) error
	ListModels(ctx context.Context, skip uint64, take uint64) ([]schema.Model, error)
}
