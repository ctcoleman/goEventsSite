package kafka

import (
	"encoding/json"
	"fmt"
	"goEventsSite/src/contracts"
	"goEventsSite/src/lib/msgqueue"
	"log"

	"github.com/Shopify/sarama"
	"github.com/mitchellh/mapstructure"
)

// kafkaEventListener defines the kafka cosumer structure
type kafkaEventListener struct {
	consumer   sarama.Consumer
	partitions []int32
}

func (ke *kafkaEventListener) Listen(events ...string) (<-chan msgqueue.Event, <-chan error, error) {
	var err error
	topic := "events"
	results := make(chan msgqueue.Event)
	errors := make(chan error)
	partitions := ke.partitions
	if len(partitions) == 0 {
		partitions, err = ke.consumer.Partitions(topic)
		if err != nil {
			return nil, nil, err
		}
	}

	log.Printf("topic %s has partitions: %v", topic, partitions)

	for _, partitions := range partitions {
		conn, err := ke.consumer.ConsumePartition(topic, partition, 0)
		if err != nil {
			return nil, nil, err
		}

		go func() {
			for msg := range conn.Messages() {
				// decode the fucking event data
				body := messageEnvelope{}
				err := json.Unmarshal(msg.Value, &body)
				if err != nil {
					errors <- fmt.Errorf("could not decode JSON message data: %s", err)
					continue
				}

				// what type of fucking event is it??
				var event msgqueue.Event
				switch body.EventName {
				case "event.created":
					event = &contracts.EventCreatedEvent{}
				case "location.created":
					event = &contracts.LocationCreatedEvent{}
				case "event.booked":
					event = &contracts.EventBookedEvent{}
				default:
					errors <- fmt.Errorf("unknown event type: %s", body.EventName)
					continue
				}

				// decode the god damn int to an actual event
				config := mapstructure.DecoderConfig{
					Result:  event,
					TagName: "json",
				}
				err = mapstructure.NewDecoder(&config).Decode(body.Payload)
				if err != nil {
					errors <- fmt.Errorf("could not map event %s: %s", body.EventName, err)
				}

				// publish that shit to the results channel
				results <- event
			}
		}()
	}
}

// NewKafkaEventListener allows you to initialize a new consumer
func NewKafkaEventListener(client sarama.Client, partitions []int32) (msgqueue.EventListener, error) {
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		return nil, err
	}

	listener := &kafkaEventListener{
		consumer:   consumer,
		partitions: partitions,
	}

	return listener, nil
}
