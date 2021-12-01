package main

import (
	"flag"
	"fmt"
	"goEventsSite/src/eventsservice/rest"
	"goEventsSite/src/lib/configuration"
	"goEventsSite/src/lib/msgqueue"
	msgqueue_amqp "goEventsSite/src/lib/msgqueue/amqp"
	"goEventsSite/src/lib/msgqueue/kafka"
	"goEventsSite/src/lib/persistence/dblayer"

	"github.com/Shopify/sarama"
	"github.com/streadway/amqp"
)

func main() {
	var eventEmitter msgqueue.EventEmitter

	confPath := flag.String("conf", `../lib/configuration/eventsconfig.json`, "flag to set path to config json file")
	flag.Parse()
	// extract the config
	config, _ := configuration.ExtractConfiguration(*confPath)

	// are we using AMQP (RabbitMQ) or is it fucking Kafka
	switch config.MessageBrokerType {
	case "amqp":
		// connect to the AMQP broker
		conn, err := amqp.Dial(config.AMQPMessageBroker)
		if err != nil {
			panic("-- error connecting to amqp broker -- " + err.Error())
		}

		// connect to the AMQP emitter
		eventEmitter, err = msgqueue_amqp.NewAMQPEventEmitter(conn, "events")
		if err != nil {
			panic("-- error connecting to amqp emitter -- " + err.Error())
		}
	case "kafka":
		// connect to kafka brokers
		saramaConfig := sarama.NewConfig()
		client, err := sarama.NewClient(config.KafkaMessageBroker, saramaConfig)
		if err != nil {
			panic(err)
		}

		// connect to kafka emitter
		eventEmitter, err = kafka.NewKafkaEventEmitter(client)
		if err != nil {
			panic(err)
		}
	default:
		panic("Bad message broker type: " + config.MessageBrokerType)
	}

	// connect to the database
	fmt.Println("Connecting to Database...")
	dbhandler, err := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	if err != nil {
		panic("-- error connecting to db ---" + err.Error())
	}
	fmt.Println("Successfully connected to Database...")

	// RestfulApi start http and https
	fmt.Println("Starting http(s) restful service router")
	httpErrChan, httptlsErrChan := rest.ServeAPI(config.RestfulEndpoint, config.RestfulTLSEndpoint, config.RestfulTLSCert, config.RestfulTLSKey, dbhandler, eventEmitter)
	select {
	case err := <-httpErrChan:
		panic("--error conencting booking rest api over http -- " + err.Error())
	case err := <-httptlsErrChan:
		panic("--error conencting booking rest api over https -- " + err.Error())
	}
}
