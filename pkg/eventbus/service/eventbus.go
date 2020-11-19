package eventbus

import (
	//"time"

	"github.com/BrobridgeOrg/gravity-exporter-kafka/pkg/app"
	log "github.com/sirupsen/logrus"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type Options struct {
	ClientName string
}

type EventBusHandler struct {
	Reconnect  func()
	Disconnect func()
}

type EventBus struct {
	app        app.App
	connection *kafka.Producer
	host       string
	handler    *EventBusHandler
	options    *Options
}

func NewEventBus(a app.App, host string, options Options) *EventBus {
	return &EventBus{
		app:        a,
		connection: nil,
		host:       host,
		options:    &options,
	}
}

func (eb *EventBus) Connect() error {

	log.WithFields(log.Fields{
		"host": eb.host,
	}).Info("Connecting to Kafka server")

	conn, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": eb.host,
	})
	if err != nil {
		log.Error(err)
	}

	eb.connection = conn

	return nil
}

func (eb *EventBus) Close() {
	eb.connection.Close()
}

func (eb *EventBus) GetConnection() *kafka.Producer {
	return eb.connection
}
