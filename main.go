package main

import (
	"fmt"
	"go-kafka/config"
	"go-kafka/consumer"
	"go-kafka/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("kafka")

	config.CreateProducer()
	config.CreateConsumer()
	go consumer.ConsumeMessage()

	router := gin.Default()
	router.POST("/publish-message", controller.PublishMessage)
	router.Run(":8080")
}
