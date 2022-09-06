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

var rep Repository

func SetRepository(repository Repository) {
	rep = repository
}

func Close() {
	rep.Close()
}

func InsertModel(ctx context.Context, model schema.Model) error {
	return rep.InsertModel(ctx, model)
}

func ListModels(ctx context.Context, skip uint64, take uint64) ([]schema.Model, error) {
	return rep.ListModels(ctx, skip, take)
}
