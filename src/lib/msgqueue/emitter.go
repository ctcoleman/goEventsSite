package msgqueue

// define interface AMQP event emitter/publisher
type EventEmitter interface {
	Emit(event Event) error
}
