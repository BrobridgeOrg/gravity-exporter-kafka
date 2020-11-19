package exporter

import (
	"io"
	"sync"
	"sync/atomic"

	"golang.org/x/net/context"

	pb "github.com/BrobridgeOrg/gravity-api/service/exporter"
	app "github.com/BrobridgeOrg/gravity-exporter-kafka/pkg/app"
	log "github.com/sirupsen/logrus"
	kafka "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

var counter uint64 = 0
var SendEventSuccess = pb.SendEventReply{
	Success: true,
}

type Event struct {
	Channel string
	Payload []byte
}

var eventPool = sync.Pool{
	New: func() interface{} {
		return &Event{}
	},
}

type Service struct {
	app      app.App
	incoming chan *Event
}

func NewService(a app.App) *Service {

	service := &Service{
		app:      a,
		incoming: make(chan *Event, 204800),
	}

	go service.eventHandler()

	return service
}

func (service *Service) eventHandler() {

	for {
		select {
		case event := <-service.incoming:
			conn := service.app.GetEventBus().GetConnection()

			msg := kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic:     &event.Channel,
					Partition: kafka.PartitionAny,
				},
				Value: event.Payload,
			}

			conn.Produce(&msg, nil)
			conn.Flush(300) //ms
		}
	}
}

func (service *Service) SendEvent(ctx context.Context, in *pb.SendEventRequest) (*pb.SendEventReply, error) {
	/*
		id := atomic.AddUint64((*uint64)(&counter), 1)

		if id%1000 == 0 {
			log.Info(id)
		}
	*/
	conn := service.app.GetEventBus().GetConnection()

	msg := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &in.Channel,
			Partition: kafka.PartitionAny,
		},
		Value: in.Payload,
	}

	err := conn.Produce(&msg, nil)
	if err != nil {
		log.Error(err)
		return &pb.SendEventReply{
			Success: false,
		}, err
	}

	conn.Flush(300) //ms

	return &SendEventSuccess, nil
}

func (service *Service) SendEventStream(stream pb.Exporter_SendEventStreamServer) error {

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}
		id := atomic.AddUint64((*uint64)(&counter), 1)

		if id%1000 == 0 {
			log.Info(id)
		}
		event := eventPool.Get().(*Event)
		event.Channel = in.Channel
		event.Payload = in.Payload

		service.incoming <- event
	}
}
