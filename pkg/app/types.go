package app

import (
	"github.com/BrobridgeOrg/gravity-exporter-kafka/pkg/eventbus"
	"github.com/BrobridgeOrg/gravity-exporter-kafka/pkg/grpc_server"
	"github.com/BrobridgeOrg/gravity-exporter-kafka/pkg/mux_manager"
)

type App interface {
	GetGRPCServer() grpc_server.Server
	GetMuxManager() mux_manager.Manager
	GetEventBus() eventbus.EventBus
}
