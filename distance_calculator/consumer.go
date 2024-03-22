package main

import (
	"encoding/json"
	"fmt"

	"github.com/adarsh-jaiss/microservice-toll-calculator/types"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct{
	consumer *kafka.Consumer
	IsRunning bool
	CalcService CalculatorServicer
}

func NewkafkaConsumer(topic string, svc CalculatorServicer) (*KafkaConsumer,error) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost",
        "group.id":          "myGroup",
        "auto.offset.reset": "earliest",
    })
    if err != nil {
        return nil, fmt.Errorf("failed to create Kafka consumer: %s", err)
    }

    err = c.SubscribeTopics([]string{topic}, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to subscribe to topic %s: %s", topic, err)
    }

    fmt.Printf("Kafka consumer subscribed to topic: %s\n", topic)

	// c.Close()

	return &KafkaConsumer{
		consumer: c,
		CalcService: svc ,
	},nil 
}

func(c *KafkaConsumer) Start() {
	logrus.Info("Kafka transport started")
	c.IsRunning = true
	c.ReadMessageLoop()
}

func (c *KafkaConsumer) ReadMessageLoop() {
	for c.IsRunning {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			logrus.Errorf("kafka consume error %s", err)
			continue
		}
		var data types.OBUData
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			logrus.Errorf("JSON serialization error: %s", err)
			continue
		}
		distance, err := c.CalcService.CalculateDistance(data)
		if err != nil {
			logrus.Errorf("calculation error: %s", err)
			continue
		}
		fmt.Printf("distance : %.2f",distance)
	}
}



	