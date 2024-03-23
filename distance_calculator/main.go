package main

import (
	"log"

	"github.com/adarsh-jaiss/microservice-toll-calculator/aggregator/client"
)

const (
	topic              = "obuData"
	aggregatorEndpoint = "http://localhost:3000/aggregate"
)

// Transport --> JSON,GRPC,Kafka --> Attach business logic to this transport
func main() {
	var (
		err error
		svc CalculatorServicer
	)

	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)

	httpClient := client.NewHTTPClient(aggregatorEndpoint)
	// grpcClient, err := client.NewGRPCClient(aggregatorEndpoint)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	KafkaConsumer, err := NewKafkaConsumer(topic, svc, httpClient)
	if err != nil {
		log.Fatal(err)
	}

	KafkaConsumer.Start()
}
