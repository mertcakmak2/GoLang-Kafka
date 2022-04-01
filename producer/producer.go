package producer

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var Producer *kafka.Producer

func ProduceMessage(topicName string, message string) error {
	kafkaEventChan := make(chan kafka.Event)
	partition := kafka.PartitionAny
	err := Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topicName, Partition: partition},
		Value:          []byte(message),
	}, kafkaEventChan)

	if err != nil {
		fmt.Println("Producer failed:", err)
	}
	d := <-kafkaEventChan
	m := d.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		fmt.Println("Failed: ", m.TopicPartition.Error)
	}
	close(kafkaEventChan)
	return m.TopicPartition.Error
}
