# Microservice-toll-calculator

## Project Overview
![Architecture drawio](https://github.com/Adarsh-jaiss/Microservice-toll-calculator/assets/96974600/18fa76f9-5a29-4b81-a333-5f0c8ee360f7)


In the main.go file, We will simulate an OBU - Onboard unit that sits in the truck and that's going to send out GPS conrdinates after each interval and we are going to send that/ replicate that using some kind of web sockets connection that will basically send these messages over webs sockets and we are gonna receive that in ou 1st microservice and put them on kafka. Now the another microservice i.e Distance calculator will use these corrdinates and calculate the distance Travelled by the vehichle and will send it to the invoicer (the another microservice) and now the Invoicer will send data to Invoice calculator to calculate and generate an Innvoive and then it will send it back to invoicer. Also this invoicer is connected to a database so it will store the generated invoice in the db and also query it back and send it to the client via the API Gateway.

Note : We have kept invoice calculator service isolated and treating it as a standalone microservice because what if someone wants to calculate the the amount of money they need to pay for travelling from place A to place B, In that case it will do the calculation and send it to user directly.


## Project Dependencies
### websoket

```
go get github.com/gorilla/websocket
```

#### [Kafka](https://github.com/confluentinc/confluent-kafka-go)

#### Kafka Go-client
```
go get github.com/confluentinc/confluent-kafka-go/v2/kafka
```

#### kafka docker installation
```
docker-compose up -d
```
### Logger

```bash
go get github.com/sirupsen/logrus
```

### Note : This project is currently in its development phase,as i had my exams, so need to pause it for a month.
