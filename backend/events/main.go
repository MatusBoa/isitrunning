package events

type EventDispatcher interface {
	Dispatch(topic string, event Event) error
}

type Event interface {
	ToString() string
}