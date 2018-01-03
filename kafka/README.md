# Kafka
         
Library kafka containing the basic function used by meastor api micro service to communicate with [kafka](http://kafka.apache.org/).


## Dependency
-  [wvanbergen/kafka](https://github.com/wvanbergen/kafka) 
-  [Shopify/sarama](https://github.com/Shopify/sarama) 

## Run The Test
In order to run the test of this library you need to have an kafka server running.
You can find a docker image of kafka [here](https://hub.docker.com/r/spotify/kafka/)
To tun 
```bash
    docker run -p 2181:2181 -p 9092:9092 --env ADVERTISED_HOST=`docker-machine ip default` --env ADVERTISED_PORT=9092  spotify/kafka
```

First set up :
- ZOOKEEPER_TEST environment variable with your zookeper server IP 
- KAFKA_TEST with you kafka server IP:
```bash
export ZOOKEEPER_TEST=$zookeeperIP:$zookeeperPort
export KAFKA_TEST=$kafkaIP:$kafkaPort
```
Then run:
```bash
go test
```

## How To Use
### Producer
Just create a new Producer and initialize it :
```go
    kafkaProducer = new(kafka.Producer)
	kafkaProducer.NewProducer(brockerList, topics, config)
```
This will create a producer  and initalise the list of topics by sendin an "init maestro" message

You can then use this interface to sendData to kafka
example :
```go
    kafkaProducer.SendData(typeRequest, "testData")
```

### Consumer
 In order to create a consumer you need to register it :
 ```go
 		kafkaConsumer, err = kafka.RegisterConsumer(&kafka.Consumer{
			Topic:        "topicName",
			Group:        "groupName",
			BrokerList:   brokerList,
			Config:       config,
			GetModel:     getModel,
			GetEventType: getEventName,
		})
```
getModel is a function that change a []byte into an interface

getEventName is a function that return the name of the event to treat depending on the model read by the consumer
```go
 func getModel(value []byte) interface{}
 func getEventName(model interface{}) string 
```
Then you just need to register an event :
```go
kafkaConsumer.RegisterEvent("issue_updated", &kafka.ProcessEvent{processUpdateData})
```
Where ProcessEvent is the function in which you treat the event:
```bash
func processUpdateData(event interface{})
```

### Example:
```go
func Kafka() *kafka.Consumer {
		kafkaConsumer, err = kafka.RegisterConsumer(&kafka.Consumer{
			Topic:        "testTopic",
			Group:        "groupTest",
			BrokerList:   brokerList,
			Config:       config,
			GetModel:     getModel,
			GetEventType: getEventName,
		})
		if err != nil {
			log.Fatal("An error occur while connecting to kafka, err : ", err)
		}
	return kafkaConsumer
}

func getModel(value []byte) interface{} {
	return string(value)
}

func getEventName(model interface{}) string {
	return "Test"
}

func init() {
   kafkaConsumer.RegisterEvent("Test", &ProcessEvent{processtest})
}

func processtest(test interface{}) {
    log.Print(test)
}
```