package kafka

import (
	"encoding/json"
	kafkahelper "goEventsSite/src/lib/helper/kafka"
	"goEventsSite/src/lib/msgqueue"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

type kafkaEventEmitter struct {
	producer sarama.SyncProducer
}

type messageEnvelope struct {
	EventName string      `json:"eventName"`
	Payload   interface{} `json:"payload"`
}

func NewKafkaEventEmitterFromEnvironment() (msgqueue.EventEmitter, error) {
	brokers := []string{"localhost:9092"}

	if brokerList := os.Getenv("KAFKA_BROKERS"); brokerList != "" {
		brokers = strings.Split(brokerList, ",")
	}

	client := <-kafkahelper.RetryConnect(brokers, 5*time.Second)
	return NewKafkaEventEmitter(client)
}

// NewKafkaEventEmitter initilizes a new Kafka emitter
func NewKafkaEventEmitter(client sarama.Client) (msgqueue.EventEmitter, error) {
	// create the emitter and connect to client
	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return nil, err
	}

	emitter := kafkaEventEmitter{
		producer: producer,
	}

	return &emitter, nil
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
		Topic:     "events",
		Key:       nil,
		Value:     sarama.ByteEncoder(jsonBody),
		Headers:   []sarama.RecordHeader{},
		Metadata:  nil,
		Offset:    0,
		Partition: 0,
		Timestamp: time.Time{},
	}

	// publish the fucking message
	_, _, err = e.producer.SendMessage(msg)
	if err != nil {
		return err
	}
	log.Printf("published message with topic %s: %v", event.EventName(), jsonBody)

	return err
}
