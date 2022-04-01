package config

import (
	"fmt"
	"go-kafka/consumer"
	"go-kafka/producer"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func CreateProducer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka:9092"})
	if err != nil {
		fmt.Println(err.Error())
	}
	producer.Producer = p
	return p
}

func CreateConsumer() *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "messageGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	consumer.Consumer = c
	return c
}
