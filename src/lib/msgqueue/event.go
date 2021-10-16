package msgqueue

// define interface for AMQP events
type Event interface {
	PartitionKey() string
	EventName() string
}
