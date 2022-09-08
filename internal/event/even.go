package event

import "github.com/Olexander753/WB_L0/internal/schema"

type EventStore interface {
	Close()
	PublishModelCreated(model schema.Model) error
	SubscribeModelCreated() (<-chan ModelCreatedMessage, error)
	OnModelCreated(f func(ModelCreatedMessage)) error
}

var event EventStore

func SetEventStore(es EventStore) {
	event = es
}

func Close() {
	event.Close()
}

func PublishModelCreated(model schema.Model) error {
	return event.PublishModelCreated(model)
}

func SubscribeModelCreated() (<-chan ModelCreatedMessage, error) {
	return event.SubscribeModelCreated()
}

func OnModelCreated(f func(ModelCreatedMessage)) error {
	return event.OnModelCreated(f)
}
