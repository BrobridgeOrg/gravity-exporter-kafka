package eventbus

import (
	//"time"

	"github.com/BrobridgeOrg/gravity-exporter-kafka/pkg/app"
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

type EventBus struct {
	app        app.App
	connection sarama.AsyncProducer
	hosts      []string
}

func NewEventBus(a app.App, hosts []string) *EventBus {
	return &EventBus{
		app:        a,
		hosts:      hosts,
		connection: nil,
	}
}

func (eb *EventBus) Connect() error {

	log.WithFields(log.Fields{
		"host": eb.hosts,
	}).Info("Connecting to Kafka server")

	config := sarama.NewConfig()

	prd, err := sarama.NewAsyncProducer(eb.hosts, config)
	if err != nil {
		return err
	}

	eb.connection = prd

	return nil
}

func (eb *EventBus) Close() {
	eb.connection.Close()
}

func (eb *EventBus) GetConnection() sarama.AsyncProducer {
	return eb.connection
}
