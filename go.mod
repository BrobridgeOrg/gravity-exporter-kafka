module github.com/BrobridgeOrg/gravity-exporter-kafka

go 1.13

require (
	github.com/BrobridgeOrg/gravity-sdk v0.0.43
	github.com/Shopify/sarama v1.29.0
	github.com/json-iterator/go v1.1.10
	github.com/prometheus/common v0.4.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.7.1
	go.uber.org/automaxprocs v1.3.0
	golang.org/x/net v0.0.0-20210427231257-85d9c07bbe3a
	gopkg.in/confluentinc/confluent-kafka-go.v1 v1.5.2
)

//replace github.com/BrobridgeOrg/gravity-api v0.0.0-20200808075207-3326e6e4eea1 => ../gravity-api
