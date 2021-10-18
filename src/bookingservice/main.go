package main

import (
	"flag"
	"fmt"
	"goEventsSite/src/bookingservice/listener"
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
	var eventListener msgqueue.EventListener

	// get user config if it exist
	confPath := flag.String("config", "./configuration/config.json", "path to config file")
	flag.Parse()
	// extract the config
	config, _ := configuration.ExtractConfiguration(*confPath)

	// are we using AMQP (RabbitMQ) or is it fucking kafka
	switch config.MessageBrokerType {
	case "amqp":
		// connect the god damn booking service to the AMQP broker
		conn, err := amqp.Dial(config.AMQPMessageBroker)
		if err != nil {
			panic(err)
		}

		// connect to the god damn AMQP listener
		eventListener, err = msgqueue_amqp.NewAMQPEventListener(conn, "events")
		if err != nil {
			panic(err)
		}

		// connect to the AMQP emitter
		eventEmitter, err = msgqueue_amqp.NewAMQPEventEmitter(conn)
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

		// connect to kafka listener
		eventListener, err = kafka.NewKafkaEventListener(client, []int32{})
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

	processor := listener.EventProcessor{eventListener, dbhandler}
	go processor.ProcessEvents()

	// RestfulApi start http and https
	httpErrChan, httptlsErrChan := rest.ServeApi(config.RestfulEndpoint, config.RestfulTLSEndpoint, config.RestfulTLSCert, config.RestfulTLSKey, dbhandler, eventEmitter)
	select {
	case err := <-httpErrChan:
		panic("--error conencting to rest over http -- " + err.Error())
	case err := <-httptlsErrChan:
		panic("--error conencting to rest over https -- " + err.Error())
	}
}
