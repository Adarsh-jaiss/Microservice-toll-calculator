package main

import (
	"fmt"
	

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct{
	consumer *kafka.Consumer
	IsRunning bool
}

func NewkafkaConsumer(topic string ) (*KafkaConsumer,error) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err!= nil{
		return nil,err
	}

	c.SubscribeTopics([]string{topic}, nil)

	// c.Close()

	return &KafkaConsumer{
		consumer: c,
	},nil 
}

func(c *KafkaConsumer) Start() {
	logrus.Info("Kafka transport started")
	c.IsRunning = true
	c.ReadMessageLoop()
}

func(c *KafkaConsumer) ReadMessageLoop()  {
	for c.IsRunning{
		msg, err := c.consumer.ReadMessage(-1)
		if err!= nil{
			logrus.Errorf("kafka consume error %s",err)
			continue
		}
		fmt.Println(msg)

	}
	
}