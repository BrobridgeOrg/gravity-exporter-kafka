package eventbus

import (
	"github.com/Shopify/sarama"
)

type EventBus interface {
	Connect() error
	Close()
	GetConnection() sarama.AsyncProducer
}
