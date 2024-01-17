package main

import (
	"fmt"
	"log"
)
const KafkaTopic = "obudata"

func main()  {
	var (
		svc CalculatorServicer
		err error
	)
	svc = NewCalculatorService()
	kafkaConsumer,err := NewkafkaConsumer(KafkaTopic,svc)
	if err!= nil{
		log.Fatal(err)
	}
	kafkaConsumer.Start()
	fmt.Println("working fine")
		
}