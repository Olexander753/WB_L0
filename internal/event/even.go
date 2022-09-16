package event

import "github.com/Olexander753/WB_L0/internal/schema"

type EventStore interface {
	Close()
	// PublishModelCreated(model schema.Model) error
	OnModelCreated(f func(schema.Model)) error
}

var event EventStore

func SetEventStore(es EventStore) {
	event = es
}

func Close() {
	event.Close()
}

// func PublishModelCreated(model schema.Model) error {
// 	return event.PublishModelCreated(model)
// }

func OnModelCreated(f func(schema.Model)) error {
	return event.OnModelCreated(f)
}
