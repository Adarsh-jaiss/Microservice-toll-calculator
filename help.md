# project setup

- Simulated the OBU data(location coordinates)
- Setup the gorilla Websocket to send that Obu data to the receiver
- Now receiver will take this data and put that on a kafka queue, so we setup the kafka
- We simply used Docker-compose yaml file to setup and run kafka and zookeeper 
