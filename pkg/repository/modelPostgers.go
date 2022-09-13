package repository

import (
	"context"
	"fmt"

	"github.com/Olexander753/WB_L0/internal/schema"
	"github.com/jmoiron/sqlx"
)

type modelPostgres struct {
	db *sqlx.DB
}

func NewModelPostgres(db *sqlx.DB) *modelPostgres {
	return &modelPostgres{db: db}
}

func (m *modelPostgres) InsertModel(ctx context.Context, model schema.Model) (string, error) {
	var uid string
	query := fmt.Sprintf("INSERT INTO %s() values($1, $2, $3) RETURNING id", modelsTable)
	row := m.db.QueryRow(query, model.Order_uid)
	if err := row.Scan(&uid); err != nil {
		return "", err
	}
	return uid, nil
}

func (m *modelPostgres) SelectModel(ctx context.Context, uid string) (schema.Model, error) {
	var model schema.Model

	query := fmt.Sprintf("SELECT * FROM %s WHERE ID = $1", modelsTable)
	err := m.db.Get(&m, query, uid)

	return model, err
}
