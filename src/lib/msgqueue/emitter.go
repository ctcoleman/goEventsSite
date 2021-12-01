package msgqueue

// define interface AMQP event emitter/publisher
type EventEmitter interface {
	Emit(e Event) error
}
