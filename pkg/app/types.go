package app

import (
	"github.com/BrobridgeOrg/gravity-exporter-kafka/pkg/eventbus"
)

type App interface {
	GetEventBus() eventbus.EventBus
}
