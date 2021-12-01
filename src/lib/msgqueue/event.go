package msgqueue

// define interface for AMQP events
type Event interface {
	EventName() string
}
