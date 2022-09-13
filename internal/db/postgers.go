package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/Olexander753/WB_L0/internal/schema"
	"github.com/Olexander753/WB_L0/pkg/config"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgres(cfg *config.Config) (*PostgresRepository, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Username, cfg.Postgres.DBName, cfg.Postgres.Password, cfg.Postgres.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{
		db: db,
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
