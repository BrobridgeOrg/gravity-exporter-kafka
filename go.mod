module github.com/BrobridgeOrg/gravity-exporter-kafka

go 1.13

require (
	github.com/BrobridgeOrg/gravity-api v0.2.2
	github.com/BrobridgeOrg/gravity-exporter-nats v0.0.0-20201103202047-d04b1add3bae
	github.com/confluentinc/confluent-kafka-go v1.5.2 // indirect
	github.com/prometheus/common v0.4.0
	github.com/sirupsen/logrus v1.6.0
	github.com/soheilhy/cmux v0.1.4
	github.com/spf13/viper v1.7.1
	go.uber.org/automaxprocs v1.3.0
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859
	google.golang.org/grpc v1.31.0
	gopkg.in/confluentinc/confluent-kafka-go.v1 v1.5.2
)

//replace github.com/BrobridgeOrg/gravity-api v0.0.0-20200808075207-3326e6e4eea1 => ../gravity-api
