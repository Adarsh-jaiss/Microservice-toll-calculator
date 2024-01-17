package main

import (
	"fmt"
	"log"
)
const kafkaTopic = "obudata"

func main()  {
	kafkaConsumer,err := NewkafkaConsumer(kafkaTopic)
	if err!= nil{
		log.Fatal(err)
	}
	kafkaConsumer.Start()
	fmt.Println("working fine")
		
}