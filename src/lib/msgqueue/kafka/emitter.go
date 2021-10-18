package kafka

import (
	"encoding/json"
	"goEventsSite/src/lib/msgqueue"
	"log"

	"github.com/Shopify/sarama"
)

type kafkaEventEmitter struct {
	producer sarama.SyncProducer
}

// Emit method sends a Kafka message to the broker
func (e *kafkaEventEmitter) Emit(event msgqueue.Event) error {
	// convert the json and envelope the message
	envelope := messageEnvelope{event.EventName(), event}
	jsonBody, err := json.Marshal(&envelope)
	if err != nil {
		return err
	}

	// prep the fucking payload
	msg := &sarama.ProducerMessage{
		Topic: event.EventName(),
		Value: sarama.ByteEncoder(jsonBody),
	}

	// publish the fucking message
	_, _, err = e.producer.SendMessage(msg)
	if err != nil {
		return err
	}
	log.Printf("published message with topic %s: %v", event.EventName(), jsonBody)

	return err
}

// NewKafkaEventEmitter initilizes a new Kafka emitter
func NewKafkaEventEmitter(client sarama.Client) (msgqueue.EventEmitter, error) {
	// create the emitter and connect to client
	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return nil, err
	}

	emitter := &kafkaEventEmitter{
		producer: producer,
	}

	return emitter, nil
}
