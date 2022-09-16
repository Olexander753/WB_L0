package event

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Olexander753/WB_L0/internal/schema"
	"github.com/Olexander753/WB_L0/pkg/repository"
	"github.com/nats-io/stan.go"
)

const (
	key         = "model.create"
	ClusterName = "test-cluster"
	ClientID    = "test-1235"
)

type NatsEventStore struct {
	nc   stan.Conn
	repo *repository.Storage
}

func NewNats(url string, repo *repository.Storage) (*NatsEventStore, error) {
	log.Println("Connect to nats")
	nc, err := stan.Connect(ClusterName, ClientID, stan.NatsURL(url))
	if err != nil {
		return nil, err
	}
	return &NatsEventStore{nc: nc,
		repo: repo}, nil
}

func (es *NatsEventStore) Close() {
	if es.nc != nil {
		es.nc.Close()
	}
}

func (es *NatsEventStore) CreateModel() (err error) {
	m := schema.Model{}
	_, err = es.nc.Subscribe(key, func(msg *stan.Msg) {
		json.Unmarshal(msg.Data, &m)
		uid, err := es.repo.InsertModel(context.Background(), m)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Insert model, order_uid =", uid)
	})
	return
}
