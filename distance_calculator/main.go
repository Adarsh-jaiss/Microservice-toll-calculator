package main

import (
	"log"
)

const topic = "obuData"

// Transport --> JSON,GRPC,Kafka --> Attach business logic to this transport
func main() {
	var (
		err error
		svc CalculatorServicer
	)

	svc = NewCalculatorService()

	KafkaConsumer, err := NewKafkaConsumer(topic, svc)
	if err != nil {
		log.Fatal(err)
	}

	KafkaConsumer.Start()
}
