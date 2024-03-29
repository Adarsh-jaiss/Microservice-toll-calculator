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

### Installing GRPC and protobuffer plugins for golang
1. Protobuffers :`go install gooogle.golang.org/protobuf/cmd/protoc-gen-go@latest`
2. GRPC : ` go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
3. Deoendencies : 
    ```
    go get google.golang.org/protobuf
    go get google.golang.org/grpc
    go get google.golang.org/genproto
    ```

### Installing protobuf compiler for linux (protoc compiler)
```
sudo apt install -y protobuf-compiler
```

#### for mac :
```
brew install protobuff
```

### Prometheus

#### Installing prometheus Go client : `go get github.com/prometheus/client_golang/prometheus`

#### docker container : 
```
docker run -d \
  --name prometheus \
  -p 127.0.0.1:9090:9090 \
  -v /path/to/your/prometheus.yml:/etc/prometheus/prometheus.yml \
  prom/prometheus

```

```
docker run -d \
  --name prometheus \
  -p 127.0.0.1:9090:9090 \
  -v /home/adarsh/myfiles/backend/Final\ projects/Microservice-toll-calculator/.config/prometheus.yml:/etc/prometheus/prometheus.yml \
  prom/prometheus

```

`docker run --name prometheus -d -p 127.0.0.1:9090:9090 prom/prometheus`

- to check if its running or not, you can access the prometheus locally on : `http://localhost:9090`
-   Use this command to run




### Note : This project is currently in its development phase,as i had my exams, so need to pause it for a month.


## How to run?

- run this command : `docker-compose up -d` to run the kafka and zookeeper on your local machine
- run these commands seperately

```
make receiver
make obu
make calculator
make agg
```
- Now you can open your thunderclient and check both the REST API's
  1. `/invoice` : This calculates the invoice of the  given OBU ID in the query
  - endpoint : `http://localhost:3000/invoice?obu=6428921451518044973`
  - METHOD : `GET`
  - Response body : 
    ```
    {
        "obuID": 6428921451518044973,
        "totalDistance": 35.86001233283358,
        "totalAmount": 112.95903884842576
    }
    ```

 2. `/aggregate` : This is used to aggregate all the data coming from distance calculator automatically, but you can also send some data as JSON if you want to do it manually here as:
    - endpoint : `http://localhost:3000/aggregate`
    - METHOD : `POST`
    - Request Body : 
        ```
        {
            "value": 20.12,
            "obuID": 1838,
            "unix": 73378
        }
        ```
    - Response (Server log):
        ```
       
        HTTP Transport running at port :3000...
        INFO[0003] aggregating distance         distance=20.12 obuid=1838 unix=73378
        INFO[0003] AggregateDistance            err="<nil>" took="47.03µs"
        
        ```
