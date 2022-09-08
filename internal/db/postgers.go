package db

import (
	"context"
	"database/sql"

	"github.com/Olexander753/WB_L0/internal/schema"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		db,
	}, nil
}

func (r *PostgresRepository) Close() {
	r.db.Close()
}

func (r *PostgresRepository) InsertModel(ctx context.Context, model schema.Model) error {
	_, err := r.db.Exec("INSERT INTO ") //TODO
	return err
}

func (r *PostgresRepository) ListModels(ctx context.Context, skip uint64, take uint64) ([]schema.Model, error) {
	rows, err := r.db.Query("SELECT FROM ") //TODO
	if err != nil {
		return nil, err
	}
	defer r.Close()

	//Parse rows
	models := []schema.Model{}
	for rows.Next() {
		model := schema.Model{}
		err := rows.Scan(&model.Order_uid) //TODO
		if err == nil {
			models = append(models, model)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return models, nil
}
