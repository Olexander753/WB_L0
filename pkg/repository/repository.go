package repository

import (
	"context"

	"github.com/Olexander753/WB_L0/internal/schema"
	"github.com/jmoiron/sqlx"
)

type Model interface {
	InsertModel(ctx context.Context, model schema.Model) (string, error)
	SelectModel(ctx context.Context, uid string) (schema.Model, error)
}

type Storage struct {
	Model
}

func NewRepository(db *sqlx.DB, ce *Cach) *Storage {
	return &Storage{
		Model: NewModelSrorage(db, ce),
	}
}
