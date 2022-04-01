package consumer

import (
	"encoding/json"
	"fmt"

	"go-kafka/model"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var Consumer *kafka.Consumer

func ConsumeMessage() {
	Consumer.SubscribeTopics([]string{"message_topic"}, nil)
	for {
		consumeMessage, err := Consumer.ReadMessage(-1)
		if err == nil {
			fmt.Println("Message:", consumeMessage.TopicPartition, string(consumeMessage.Value))
			var user model.Message
			err = json.Unmarshal(consumeMessage.Value, &user)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println("Consume message error: ", err, consumeMessage)
		}
	}
}
