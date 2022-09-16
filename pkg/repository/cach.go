package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/Olexander753/WB_L0/internal/schema"
	"github.com/jmoiron/sqlx"
)

type Cach struct {
	Models map[string]schema.Model
	db     *sqlx.DB
}

var once sync.Once

func NewCach(db *sqlx.DB) *Cach {
	log.Println("Create cach")
	models := make(map[string]schema.Model)
	return &Cach{db: db, Models: models}
}

func (c *Cach) GetModelByOrder_uid(order_uid string) (schema.Model, error) {
	model, ok := c.Models[order_uid]
	if !ok {
		err := fmt.Errorf("not found model whith order_uid: %s", order_uid)
		return model, err
	}
	return model, nil

}

func (c *Cach) SelectModels() {
	once.Do(func() {

		var mod []struct {
			Order_uid string `json:"order_uid" db:"order_uid"`
			Body      string `json:"body" db:"body"`
		}
		var model schema.Model

		query := fmt.Sprintf("SELECT * FROM %s", "Model")
		err := c.db.Select(&mod, query)
		if err != nil {
			log.Fatal(err)
		}

		for _, val := range mod {
			err = json.Unmarshal([]byte(val.Body), &model)
			if err != nil {
				log.Fatal(err)
			}
			c.Models[val.Order_uid] = model
		}
	})
}
