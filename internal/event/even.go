package event

type EventStore interface {
	Close()
	CreateModel() error
}

var event EventStore

func SetEventStore(es EventStore) {
	event = es
}

func Close() {
	event.Close()
}

func CreatedModel() error {
	return event.CreateModel()
}
