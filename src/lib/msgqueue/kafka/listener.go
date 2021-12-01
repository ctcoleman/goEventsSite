package kafka

import (
	"encoding/json"
	"fmt"
	kafkahelper "goEventsSite/src/lib/helper/kafka"
	"goEventsSite/src/lib/msgqueue"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

// kafkaEventListener defines the kafka cosumer structure
type kafkaEventListener struct {
	consumer   sarama.Consumer
	partitions []int32
	mapper     msgqueue.EventMapper
}

func NewKafkaEventListenerFromEnvironment() (msgqueue.EventListener, error) {
	brokers := []string{"localhost:9092"}
	partitions := []int32{}

	if brokerList := os.Getenv("KAFKA_BROKERS"); brokerList != "" {
		brokers = strings.Split(brokerList, ",")
	}

	if partitionList := os.Getenv("KAFKA_PARTITIONS"); partitionList != "" {
		partitionStrings := strings.Split(partitionList, ",")
		partitions = make([]int32, len(partitionStrings))

		for i := range partitionStrings {
			partition, err := strconv.Atoi(partitionStrings[i])
			if err != nil {
				return nil, err
			}
			partitions[i] = int32(partition)
		}
	}

	client := <-kafkahelper.RetryConnect(brokers, 5*time.Second)

	return NewKafkaEventListener(client, partitions)
}

// NewKafkaEventListener initializes a new Kafka consumer
func NewKafkaEventListener(client sarama.Client, partitions []int32) (msgqueue.EventListener, error) {
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		return nil, err
	}

	listener := &kafkaEventListener{
		consumer:   consumer,
		partitions: partitions,
		mapper:     msgqueue.NewEventMapper(),
	}

	return listener, nil
}

// Listen method listens for Kafka messages on the defined kafkaEventlistener
func (ke *kafkaEventListener) Listen(events ...string) (<-chan msgqueue.Event, <-chan error, error) {
	// define the topic and listening goroutine channels
	var err error
	topic := "events"
	results := make(chan msgqueue.Event)
	errors := make(chan error)

	// define the Kafka partitions for the topic
	partitions := ke.partitions
	if len(partitions) == 0 {
		partitions, err = ke.consumer.Partitions(topic)
		if err != nil {
			return nil, nil, err
		}
	}
	log.Printf("topic %s has partitions: %v", topic, partitions)

	// for each partition run the listening goroutines
	for _, partition := range partitions {
		log.Printf("consuming partition %s: %d", topic, partition)

		conn, err := ke.consumer.ConsumePartition(topic, partition, 0)
		if err != nil {
			return nil, nil, err
		}
		// listen for error and Kafka messages published to the partition
		go func() {
			for msg := range conn.Messages() {
				// decode the fucking event data
				body := messageEnvelope{}
				err := json.Unmarshal(msg.Value, &body)
				if err != nil {
					errors <- fmt.Errorf("could not decode JSON message data: %s", err)
					continue
				}

				event, err := ke.mapper.MapEvent(body.EventName, body.Payload)
				if err != nil {
					errors <- fmt.Errorf("could not map message: %v", err)
					continue
				}

				// publish that shit to the results channel
				results <- event
			}
		}()

		go func() {
			for err := range conn.Errors() {
				errors <- err
			}
		}()
	}
	return results, errors, nil
}

func (l *kafkaEventListener) Mapper() msgqueue.EventMapper {
	return l.mapper
}
