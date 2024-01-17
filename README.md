# Microservice-toll-calculator

OBU :
In the main.go file, We will simulate an
OBU - Onboard unit that sits in the truck and that's going to send out GPS conrdinates after each interval and we are going to send that/ replicate that using some kind of web sockets connection that will basically send these messages over webs sockets and we are gonna receive that in ou 1st microservice and put them on kafka   

# websoket

```
go get github.com/gorilla/websocket
```

# [Kafka](https://github.com/confluentinc/confluent-kafka-go)

### Kafka Go-client
```
go get github.com/confluentinc/confluent-kafka-go/v2/kafka
```

### kafka docker installation
```
docker-compose up -d
```
# Logger

```bash
go get github.com/sirupsen/logrus
