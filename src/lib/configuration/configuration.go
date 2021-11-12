package configuration

import (
	"encoding/json"
	"fmt"
	"goEventsSite/src/lib/persistence/dblayer"
	"os"
	"strings"
)

var (
	DBTypeDefault             = dblayer.DBTYPE("mongodb")
	DBConnectionDefault       = "mongodb://localhost:27017"
	RestfulEPDefault          = "localhost:8888"
	RestfulTLSEPDefault       = "localhost:9999"
	RestfulTLSCertDefault     = "../../etc/keys/cert.pem"
	RestfulTLSKeyDefault      = "../../etc/keys/key.pem"
	MessageBrokerTypeDefault  = "amqp"
	AMQPMessageBrokerDefault  = "amqp://guest:guest@localhost:5672"
	KafkaMessageBrokerDefault = []string{"localhost:9092"}
)

type ServiceConfig struct {
	Databasetype       dblayer.DBTYPE `json:"databasetype"`
	DBConnection       string         `json:"dbconnection"`
	RestfulEndpoint    string         `json:"restfulapi_endpoint"`
	RestfulTLSEndpoint string         `json:"restfulapi_tlsendpoint"`
	RestfulTLSCert     string         `json:"restfulapi_tlscert"`
	RestfulTLSKey      string         `json:"restfulapi_tlskey"`
	MessageBrokerType  string         `json:"message_broker_type"`
	AMQPMessageBroker  string         `json:"amqp_message_broker"`
	KafkaMessageBroker []string       `json:"kafka_message_broker"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestfulEPDefault,
		RestfulTLSEPDefault,
		RestfulTLSCertDefault,
		RestfulTLSKeyDefault,
		MessageBrokerTypeDefault,
		AMQPMessageBrokerDefault,
		KafkaMessageBrokerDefault,
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found. Using default values")
		return conf, err
	}

	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		panic("could not decode json config file" + err.Error())
	}

	if v := os.Getenv("LISTEN_URL"); v != "" {
		conf.RestfulEndpoint = v
	}

	if v := os.Getenv("MONGO_URL"); v != "" {
		conf.Databasetype = "mongodb"
		conf.DBConnection = v
	}

	if v := os.Getenv("AMQP_URL"); v != "" {
		conf.MessageBrokerType = "amqp"
		conf.AMQPMessageBroker = v
	} else if v := os.Getenv("KAFKA_BROKER_URLS"); v != "" {
		conf.MessageBrokerType = "kafka"
		conf.KafkaMessageBroker = strings.Split(v, ",")
	}

	return conf, nil
}
