package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Olexander753/WB_L0/internal/schema"
	"github.com/jmoiron/sqlx"
)

type modelStorage struct {
	db *sqlx.DB
	ce *Cach
}

func NewModelSrorage(db *sqlx.DB, ce *Cach) *modelStorage {
	return &modelStorage{
		db: db,
		ce: ce}
}

func (m *modelStorage) InsertModel(ctx context.Context, model schema.Model) (string, error) {
	var uid string
	insertModelOrder_uid := model.Order_uid
	model.Order_uid = ""
	b, err := json.Marshal(model)
	if err != nil {
		return "", err
	}

	query := fmt.Sprintf("INSERT INTO %s values($1, $2) RETURNING order_uid ;", "Model")
	row := m.db.QueryRow(query, insertModelOrder_uid, b)
	if err := row.Scan(&uid); err != nil {
		return "", err
	}

	m.ce.Models[insertModelOrder_uid] = model
	return uid, nil
}

func (m *modelStorage) SelectModel(ctx context.Context, order_uid string) (schema.Model, error) {
	model, err := m.ce.GetModelByOrder_uid(order_uid)
	return model, err
}
