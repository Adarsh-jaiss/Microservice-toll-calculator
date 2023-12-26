package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adarsh-jaiss/microservice-toll-calculator/types"
	"github.com/gorilla/websocket"
)

type DataReceiver struct{
	msgch chan types.OBUData
	conn *websocket.Conn
}

func NewDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgch: make(chan types.OBUData,128),
	}
}

func main()  {
	fmt.Println("------- Starting Data receiver --------")
	recv := NewDataReceiver()
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
		fmt.Printf("received OBUdata from [%d] :: <lat %.2f,long %.2f> \n",data.OBUID,data.Latitiude,data.Longitude)
		// dr.msgch <- data
		
		
	}
}
