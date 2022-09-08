package event

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/Olexander753/WB_L0/internal/schema"
	"github.com/nats-io/nats.go"
)

type NatsEventStore struct {
	nc                       *nats.Conn
	modelCreatedSubscription *nats.Subscription
	modelCreatedChan         chan ModelCreatedMessage
}

func NewNats(url string) (*NatsEventStore, error) {
	nc, err := nats.Connect(url)
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
	close(es.modelCreatedChan)
}

func (es *NatsEventStore) PublishModelCreated(model schema.Model) error {
	m := ModelCreatedMessage{} //TODO
	data, err := es.writeMessage(&m)
	if err != nil {
		return err
	}
	return es.nc.Publish(m.Key(), data)
}

func (es *NatsEventStore) SubscribeModelCreated() (<-chan ModelCreatedMessage, error) {
	m := ModelCreatedMessage{}
	es.modelCreatedChan = make(chan ModelCreatedMessage, 64)
	ch := make(chan *nats.Msg, 64)
	var err error
	es.modelCreatedSubscription, err = es.nc.ChanSubscribe(m.Key(), ch)
	if err != nil {
		return nil, err
	}
	// Decode message
	go func() {
		for {
			select {
			case msg := <-ch:
				if err := es.readMessage(msg.Data, &m); err != nil {
					log.Fatal(err)
				}
				es.modelCreatedChan <- m
			}
		}
	}()
	return (<-chan ModelCreatedMessage)(es.modelCreatedChan), nil
}

func (es *NatsEventStore) OnModelCreated(f func(ModelCreatedMessage)) (err error) {
	m := ModelCreatedMessage{}
	es.modelCreatedSubscription, err = es.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		es.readMessage(msg.Data, &m)
		f(m)
	})
	return
}

func (mq *NatsEventStore) writeMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (mq *NatsEventStore) readMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}
