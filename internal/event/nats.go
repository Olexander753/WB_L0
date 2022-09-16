package event

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Olexander753/WB_L0/internal/schema"
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
	//modelCreatedChan         chan schema.Model
}

func NewNats(url string) (*NatsEventStore, error) {
	log.Println("Connect to nats")
	nc, err := stan.Connect(ClusterName, ClientID, stan.NatsURL(url))
	if err != nil {
		return nil, err
	}
	return &NatsEventStore{nc: nc}, nil
}

func (es *NatsEventStore) Close() {
	if es.nc != nil {
		es.nc.Close()
	}
	if es.modelCreatedSubscription != nil {
		es.modelCreatedSubscription.Unsubscribe()
	}
	//close(es.modelCreatedChan)
}

func (es *NatsEventStore) OnModelCreated(f func(schema.Model)) (err error) {
	m := schema.Model{}
	_, err = es.nc.Subscribe(key, func(msg *stan.Msg) {
		fmt.Println(string(msg.Data))
		json.Unmarshal(msg.Data, &m)
		//es.readMessage(msg.Data, &m)
		f(m)
	})
	//fmt.Println(m)
	return
}

func ModelCreated(m schema.Model) {
	fmt.Println(m)

}

// func (mq *NatsEventStore) readMessage(data []byte, m interface{}) error {
// 	b := bytes.Buffer{}
// 	b.Write(data)
// 	return gob.NewDecoder(&b).Decode(m)
// }

//	func (es *NatsEventStore) PublishModelCreated(model schema.Model) error {
//		m := ModelCreatedMessage{} //TODO
//		data, err := es.writeMessage(&m)
//		if err != nil {
//			return err
//		}
//		return es.nc.Publish(m.Key(), data)
//	}
//
//	func (mq *NatsEventStore) writeMessage(m Message) ([]byte, error) {
//		b := bytes.Buffer{}
//		err := gob.NewEncoder(&b).Encode(m)
//		if err != nil {
//			return nil, err
//		}
//		return b.Bytes(), nil
//	}
