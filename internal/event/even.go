package event

type EventStore interface {
	Close()
	// PublishModelCreated(model schema.Model) error
	OnModelCreated() error
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

func OnModelCreated() error {
	return event.OnModelCreated()
}
