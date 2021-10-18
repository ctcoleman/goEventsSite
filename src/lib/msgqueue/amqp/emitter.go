package amqp

import (
	"encoding/json"
	"goEventsSite/src/lib/msgqueue"

	"github.com/streadway/amqp"
)

type amqpEventEmitter struct {
	connection *amqp.Connection
}

// setup method sets up a channel to the broker and defines message types
func (a *amqpEventEmitter) setup() error {
	// create the fucking AMQP channel
	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close() // always close open doors

	return channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)
}

// Emit method sends an AMQP message to the vbroker
func (a *amqpEventEmitter) Emit(event msgqueue.Event) error {
	// encode the fucking json data
	jsonDoc, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// init the fucking channel
	channel, err := a.connection.Channel()
	if err != nil {
		return nil
	}
	defer channel.Close()

	// define the AMQP message
	msg := amqp.Publishing{
		Headers:     amqp.Table{"x-event-name": event.EventName()},
		Body:        jsonDoc,
		ContentType: "application/json",
	}

	// emit that shit to the broker using defined channel
	return channel.Publish(
		"events",
		event.EventName(),
		false,
		false,
		msg,
	)
}

// NewAMQPEventEmitter initializes a new AMQP publisher
func NewAMQPEventEmitter(conn *amqp.Connection) (msgqueue.EventEmitter, error) {
	emitter := &amqpEventEmitter{
		connection: conn,
	}

	err := emitter.setup()
	if err != nil {
		return nil, err
	}

	return emitter, nil
}
