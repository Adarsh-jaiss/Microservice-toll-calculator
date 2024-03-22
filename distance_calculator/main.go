package main

import (
	
	"log"

	// "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// type DistanceCalculator struct {
// 	consumer DataConsumer
// }

const topic = "obuData"

func main() {
	KafkaConsumer, err := NewKafkaConsumer(topic)
	if err != nil {
		log.Fatal(err)
	}	
	KafkaConsumer.Start()
}
