package repository

import (
	"context"

	"github.com/Olexander753/WB_L0/internal/cach"
	"github.com/Olexander753/WB_L0/internal/schema"
	"github.com/jmoiron/sqlx"
)

type Model interface {
	InsertModel(ctx context.Context, model schema.Model) (string, error)
	SelectModel(ctx context.Context, uid string) (schema.Model, error)
}

type Repository struct {
	Model
}

func NewRepository(db *sqlx.DB, ce *cach.Cach) *Repository {
	return &Repository{
		Model: NewModelPostgres(db, ce),
	}
}
