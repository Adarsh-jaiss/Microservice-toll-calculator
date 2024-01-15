# Microservice-toll-calculator

OBU :
In the main.go file, We will simulate an
OBU - Onboard unit that sits in the truck and that's going to send out GPS conrdinates after each interval and we are going to send that/ replicate that using some kind of web sockets connection that will basically send these messages over webs sockets and we are gonna receive that in ou 1st microservice and put them on kafka   

# websoket

```
go get github.com/gorilla/websocket
```

# docker 
```
docker-compose up -d
```