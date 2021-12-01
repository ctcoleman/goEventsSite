package amqp

import (
	"encoding/json"
	"fmt"
	amqphelper "goEventsSite/src/lib/helper/amqp"
	"goEventsSite/src/lib/msgqueue"
	"os"
	"time"

	"github.com/streadway/amqp"
)

type amqpEventEmitter struct {
	connection *amqp.Connection
	exchange   string
	events     chan *emittedEvent
}

type emittedEvent struct {
	event     msgqueue.Event
	errorChan chan error
}

func NewAMQPEventEmitterFromEnvironment() (msgqueue.EventEmitter, error) {
	var url string
	var exchange string

	if url = os.Getenv("AMQP_URL"); url == "" {
		url = "amqp://localhost:5672"
	}

	if exchange = os.Getenv("AMQP_EXCHANGE"); exchange == "" {
		exchange = "example"
	}

	conn := <-amqphelper.RetryConnect(url, 5*time.Second)
	return NewAMQPEventEmitter(conn, exchange)
}

// NewAMQPEventEmitter initializes a new AMQP publisher
func NewAMQPEventEmitter(conn *amqp.Connection, exchange string) (msgqueue.EventEmitter, error) {
	emitter := amqpEventEmitter{
		connection: conn,
		exchange:   exchange,
	}

	err := emitter.setup()
	if err != nil {
		return nil, err
	}

	return &emitter, nil
}

// setup method sets up a channel to the broker and defines message types
func (a *amqpEventEmitter) setup() error {
	// create the fucking AMQP channel
	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close() // always close open doors

	err = channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)
	return err
}

// Emit method sends an AMQP message to the vbroker
func (a *amqpEventEmitter) Emit(event msgqueue.Event) error {
	// init the fucking channel
	channel, err := a.connection.Channel()
	if err != nil {
		return nil
	}
	defer channel.Close()

	// encode the fucking json data
	jsonDoc, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("could not JSON-serialize event: %s", err)
	}

	// define the AMQP message
	msg := amqp.Publishing{
		Headers:     amqp.Table{"x-event-name": event.EventName()},
		Body:        jsonDoc,
		ContentType: "application/json",
	}

	// emit that shit to the broker using defined channel
	err = channel.Publish(
		a.exchange,
		event.EventName(),
		false,
		false,
		msg,
	)

	return err
}
