package instance

import (
	"runtime"
	"strings"

	eventbus "github.com/BrobridgeOrg/gravity-exporter-kafka/pkg/eventbus/service"
	subscriber "github.com/BrobridgeOrg/gravity-exporter-kafka/pkg/subscriber/service"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AppInstance struct {
	done       chan bool
	eventBus   *eventbus.EventBus
	subscriber *subscriber.Subscriber
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

	// get kafka host
	kafkaHostStr := viper.GetString("kafka.hosts")
	kafkaHosts := strings.Split(kafkaHostStr, ",")

	// Initializing modules
	a.eventBus = eventbus.NewEventBus(a, kafkaHosts)

	a.subscriber = subscriber.NewSubscriber(a)

	// Initializing EventBus
	err := a.initEventBus()
	if err != nil {
		return err
	}

	err = a.subscriber.Init()
	if err != nil {
		return err
	}

	return nil
}

func (a *AppInstance) Uninit() {
}

func (a *AppInstance) Run() error {

	err := a.subscriber.Run()
	if err != nil {
		return err
	}

	<-a.done

	return nil
}
