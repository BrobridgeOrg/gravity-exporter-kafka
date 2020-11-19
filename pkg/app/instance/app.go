package instance

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	//	"time"

	eventbus "github.com/BrobridgeOrg/gravity-exporter-kafka/pkg/eventbus/service"
	grpc_server "github.com/BrobridgeOrg/gravity-exporter-kafka/pkg/grpc_server/server"
	mux_manager "github.com/BrobridgeOrg/gravity-exporter-kafka/pkg/mux_manager/manager"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AppInstance struct {
	done       chan bool
	muxManager *mux_manager.MuxManager
	grpcServer *grpc_server.Server
	eventBus   *eventbus.EventBus
}

func NewAppInstance() *AppInstance {

	a := &AppInstance{
		done: make(chan bool),
	}

	return a
}

func (a *AppInstance) Init() error {

	log.WithFields(log.Fields{
		"max_procs": runtime.GOMAXPROCS(0),
	}).Info("Starting application")

	// Using hostname (pod name) by default
	hostname, err := os.Hostname()
	if err != nil {
		log.Error(err)
		return err
	}

	hostname = strings.ReplaceAll(hostname, ".", "_")
	clientID := fmt.Sprintf("gravity_exporter_kafka-%s", hostname)

	// get kafka host
	kHost := viper.GetString("kafka.host")

	// Initializing modules
	a.muxManager = mux_manager.NewMuxManager(a)
	a.grpcServer = grpc_server.NewServer(a)
	a.eventBus = eventbus.NewEventBus(
		a,
		kHost,
		eventbus.Options{
			ClientName: clientID,
		},
	)

	a.initMuxManager()

	// Initializing EventBus
	err = a.initEventBus()
	if err != nil {
		return err
	}

	// Initializing GRPC server
	err = a.initGRPCServer()
	if err != nil {
		return err
	}

	return nil
}

func (a *AppInstance) Uninit() {
}

func (a *AppInstance) Run() error {

	// GRPC
	go func() {
		err := a.runGRPCServer()
		if err != nil {
			log.Error(err)
		}
	}()

	err := a.runMuxManager()
	if err != nil {
		return err
	}

	<-a.done

	return nil
}
