package event

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Olexander753/WB_L0/internal/schema"
	"github.com/Olexander753/WB_L0/pkg/repository"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

const (
	key         = "model.create"
	ClusterName = "test-cluster"
	ClientID    = "test-1235"
)

type NatsEventStore struct {
	nc                       stan.Conn
	modelCreatedSubscription *nats.Subscription
	repo                     *repository.Repository
}

func NewNats(url string, repo *repository.Repository) (*NatsEventStore, error) {
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
	if es.modelCreatedSubscription != nil {
		es.modelCreatedSubscription.Unsubscribe()
	}
}

func (es *NatsEventStore) OnModelCreated() (err error) {
	m := schema.Model{}
	_, err = es.nc.Subscribe(key, func(msg *stan.Msg) {
		json.Unmarshal(msg.Data, &m)
		//fmt.Println(m)
		_, err := es.repo.InsertModel(context.Background(), m)
		if err != nil {
			log.Fatal(err)
		}
	})
	return
}
