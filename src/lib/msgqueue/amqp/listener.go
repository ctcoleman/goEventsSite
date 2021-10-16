package amqp

import (
	"encoding/json"
	"fmt"
	"goEventsSite/src/contracts"
	"goEventsSite/src/lib/msgqueue"

	"github.com/streadway/amqp"
)

// amqpEventListener defines the connection and queue to listen for AMQP events
type amqpEventListener struct {
	connection *amqp.Connection
	queue      string
}

// setup method creates a channel to the AMQP broker and defines the queue to listen on
func (a *amqpEventListener) setup() error {
	// create a a god damn channel to the AMQP broker
	channel, err := a.connection.Channel()
	if err != nil {
		return nil
	}
	defer channel.Close()

	// declare the fucking AMQP Queue to listen on....you know....queue fucking messaging??
	_, err = channel.QueueDeclare(a.queue, true, false, false, false, nil)

	return err
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
		if err := channel.QueueBind(a.queue, eventName, "events", false, nil); err != nil {
			return nil, nil, err
		}
	}

	// consume message events from the queue. nom...nom...nom
	msgs, err := channel.Consume(a.queue, "", false, false, false, false, nil)
	if err != nil {
		return nil, nil, err
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

			// create a new fucking event
			var event msgqueue.Event

			switch eventName {
			case "eventCreated":
				event = new(contracts.EventCreatedEvent)
			default:
				errors <- fmt.Errorf("event type %s is unknown", eventName)
				continue
			}

			// store the message body json data
			err := json.Unmarshal(msg.Body, event)
			if err != nil {
				errors <- err
				continue
			}
			events <- event
		}
	}()

	return events, errors, nil
}

// NewAMQPEventListener define new AMQP listener and pass it through setup
func NewAMQPEventListener(conn *amqp.Connection, queue string) (msgqueue.EventListener, error) {
	// define the listener with the given connection and queue details
	listener := &amqpEventListener{
		connection: conn,
		queue:      queue,
	}

	// run that shit through the setup function
	err := listener.setup()
	if err != nil {
		return nil, err
	}

	return listener, nil
}
