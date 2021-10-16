package msgqueue

// define interface for AMQP event listener
type EventListener interface {
	Listen(eventNames ...string) (<-chan Event, <-chan error, error)
}
