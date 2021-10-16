package main

import (
	"flag"
	"fmt"
	"goEventsSite/src/eventsservice/rest"
	"goEventsSite/src/lib/configuration"
	msgqueue_amqp "goEventsSite/src/lib/msgqueue/amqp"
	"goEventsSite/src/lib/persistence/dblayer"

	"github.com/streadway/amqp"
)

func main() {
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set path to config json file")
	flag.Parse()

	// extract the config
	config, _ := configuration.ExtractConfiguration(*confPath)

	// connect to the AMQP broker
	conn, err := amqp.Dial(config.AMQPMessageBroker)
	if err != nil {
		panic("-- error connecting to amqp broker -- " + err.Error())
	}

	// connect to the AMQP emitter
	eventEmitter, err := msgqueue_amqp.NewAMQPEventEmitter(conn)
	if err != nil {
		panic("-- error connecting to amqp emitter -- " + err.Error())
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
	httpErrChan, httptlsErrChan := rest.ServeApi(config.RestfulEndpoint, config.RestfulTLSEndpoint, config.RestfulTLSCert, config.RestfulTLSKey, dbhandler, eventEmitter)
	select {
	case err := <-httpErrChan:
		panic(err)
	case err := <-httptlsErrChan:
		panic(err)
	}
}
