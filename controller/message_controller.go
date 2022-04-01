package controller

import (
	"encoding/json"
	"go-kafka/model"
	"go-kafka/producer"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PublishMessage(c *gin.Context) {

	var message model.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message.ID = strconv.Itoa(rand.Intn(200))
	data, err := json.Marshal(message)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed when converting")
	}
	producer.ProduceMessage("message_topic", string(data))
	c.String(http.StatusOK, "published")
}
