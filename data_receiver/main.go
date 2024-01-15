package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/adarsh-jaiss/microservice-toll-calculator/types"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gorilla/websocket"
)

var KafkaTopic = "obudata"

type DataReceiver struct {
    msgch chan types.OBUData
    conn  *websocket.Conn
    Prod  *kafka.Producer
}

func NewDataReceiver() (*DataReceiver,error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		return nil,err
	}
		// Start another go routine to check if we have delivered the data
		go func() {
			for e := range p.Events() {
				switch ev := e.(type) {
				case *kafka.Message:
					if ev.TopicPartition.Error != nil {
						fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
					} else {
						fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
					}
				}
			}
		}()

	return &DataReceiver{
		msgch: make(chan types.OBUData,128),
		Prod: p,
	},nil
}

func (dr *DataReceiver) DataProducer(data types.OBUData) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = dr.Prod.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &KafkaTopic,
			Partition: kafka.PartitionAny,
		},
		Value: b,
	}, nil)

	return err
}


func main()  {

	fmt.Println("------- Starting Data receiver --------")
	recv,err := NewDataReceiver()
	if err!=nil{
		log.Fatal(err)
	}
	http.HandleFunc("/ws", recv.HandleWS)
	http.ListenAndServe(":30000",nil)

}

func(dr *DataReceiver) HandleWS(w http.ResponseWriter, r *http.Request)  {
	 u:= websocket.Upgrader{
		ReadBufferSize: 1028,
		WriteBufferSize: 1028,
	 }
	conn,err := u.Upgrade(w,r, nil)
	if err!=nil{
		log.Fatal(err)
	}
	dr.conn = conn
	
	go dr.wsReceiverLoop()
}

func (dr *DataReceiver) wsReceiverLoop()  {
	fmt.Println("New OBU connected Client connected")
	defer func() {
        fmt.Println("OBU Client disconnected")
        close(dr.msgch)
    }()

	
	for{
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err!= nil{
			log.Println("read error :",err)
			continue
			
		}
		// fmt.Println("received message : ",data)
		if err := dr.DataProducer(data); err!= nil{
			fmt.Println("kafka produce error :", err)
		}
		// fmt.Printf("received OBUdata from [%d] :: <lat %.2f,long %.2f> \n",data.OBUID,data.Latitiude,data.Longitude)
		// dr.msgch <- data
		
		
	}
}
