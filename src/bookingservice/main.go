package main

import (
	"flag"
	"goEventsSite/src/bookingservice/listener"
	"goEventsSite/src/eventsservice/rest"
	"goEventsSite/src/lib/configuration"
	"goEventsSite/src/lib/msgqueue"
	msgqueue_amqp "goEventsSite/src/lib/msgqueue/amqp"
	"goEventsSite/src/lib/persistence"
	"goEventsSite/src/lib/persistence/dblayer"

	"github.com/streadway/amqp"
)

func main() {
	var eventListener msgqueue.EventListener
	var dbhandler persistence.DatabaseHandler

	// get user config if it exist
	confPath := flag.String("config", "./configuration/config.json", "path to config file")
	flag.Parse()
	config, _ := configuration.ExtractConfiguration(*confPath)

	// connect the fucking booking service to the database
	dbhandler, err := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	if err != nil {
		panic(err)
	}

	// connect the god damn booking service to the AMQP broker
	conn, err := amqp.Dial(config.AMQPMessageBroker)
	if err != nil {
		panic(err)
	}

	// connect to the AMQP emitter
	eventEmitter, err := msgqueue_amqp.NewAMQPEventEmitter(conn)
	if err != nil {
		panic("-- error connecting to amqp emitter -- " + err.Error())
	}

	// create a new god damn event listener
	eventListener, err = msgqueue_amqp.NewAMQPEventListener(conn, "events")
	if err != nil {
		panic(err)
	}

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
