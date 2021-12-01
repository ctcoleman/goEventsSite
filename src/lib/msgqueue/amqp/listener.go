package amqp

import (
	"fmt"
	"goEventsSite/src/lib/msgqueue"
	"os"
	"time"

	amqphelper "goEventsSite/src/lib/helper/amqp"

	"github.com/streadway/amqp"
)

// amqpEventListener defines the connection and queue to listen for AMQP events
type amqpEventListener struct {
	connection *amqp.Connection
	exchange   string
	queue      string
	mapper     msgqueue.EventMapper
}

func NewAMQPEventListenerFromEnvironment() (msgqueue.EventListener, error) {
	var url string
	var exchange string
	var queue string

	if url = os.Getenv("AMQP_URL"); url == "" {
		url = "amqp://localhost:5672"
	}

	if exchange = os.Getenv("AMQP_EXCHANGE"); exchange == "" {
		exchange = "example"
	}

	if exchange = os.Getenv("AMQP_QUEUE"); queue == "" {
		queue = "example"
	}

	conn := <-amqphelper.RetryConnect(url, 5*time.Second)
	return NewAMQPEventListener(conn, exchange, queue)
}

// NewAMQPEventListener define new AMQP listener and pass it through setup
func NewAMQPEventListener(conn *amqp.Connection, exchange string, queue string) (msgqueue.EventListener, error) {
	listener := amqpEventListener{
		connection: conn,
		exchange:   exchange,
		queue:      queue,
		mapper:     msgqueue.NewEventMapper(),
	}

	err := listener.setup()
	if err != nil {
		return nil, err
	}

	return &listener, nil
}

// Listen function listens for events based on given name and maps those events to the corresponding event structure
func (a *amqpEventListener) Listen(eventNames ...string) (<-chan msgqueue.Event, <-chan error, error) {
	// create a fucking connection to the channel
	channel, err := a.connection.Channel()
	if err != nil {
		return nil, nil, err
	}
	defer channel.Close() // always close open doors....

	for _, eventName := range eventNames {
		// bind the event names to the given fucking queue so we can listen for the right fucking events
		if err := channel.QueueBind(a.queue, eventName, a.exchange, false, nil); err != nil {
			return nil, nil, fmt.Errorf("could not bind event %s to queue %s: %s", eventName, a.queue, err)
		}
	}

	// consume message events from the queue. nom...nom...nom
	msgs, err := channel.Consume(a.queue, "", false, false, false, false, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("could not consume queue: %s", err)
	}

	// create events and error variables from each channel
	events := make(chan msgqueue.Event)
	errors := make(chan error)

	go func() {
		// map the god damn messages to their  actual fucking message structures
		for msg := range msgs {
			// ensure there even is a fucking event-name header on the message
			rawEventName, ok := msg.Headers["x-event-name"]
			if !ok {
				errors <- fmt.Errorf("msg did not contain x-event-name header")
				msg.Nack(false, false) // negative acknowledgment - denied bitch!
				continue
			}

			// stringify that shit
			eventName, ok := rawEventName.(string)
			if !ok {
				errors <- fmt.Errorf(
					"x-event-name header is not string, but %t",
					rawEventName,
				)
				msg.Nack(false, false) // negative acknowledgment - denied bitch!
				continue
			}

			event, err := a.mapper.MapEvent(eventName, msg.Body)
			if err != nil {
				errors <- fmt.Errorf("could not unmarshal event %s: %s", eventName, err)
				msg.Nack(false, false)
				continue
			}

			events <- event
			msg.Ack(false)
		}
	}()

	return events, errors, nil
}

func (l *amqpEventListener) Mapper() msgqueue.EventMapper {
	return l.mapper
}

// setup method creates a channel to the AMQP broker and defines the queue to listen on
func (a *amqpEventListener) setup() error {
	channel, err := a.connection.Channel()
	if err != nil {
		return nil
	}
	defer channel.Close()

	err = channel.ExchangeDeclare(a.exchange, "topic", true, false, false, false, nil)
	if err != nil {
		return err
	}

	_, err = channel.QueueDeclare(a.queue, true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("could not declare queue %s: %s", a.queue, err)
	}

	return nil
}
