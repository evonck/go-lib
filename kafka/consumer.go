/*
Package kafka provides a client libraries to communicate with kafka. The Consumer
object is the high-level API for consuming message.
The Producer object is the high-levle API for producing message.
*/
package kafka

import (
	"errors"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
	"github.com/wvanbergen/kafka/consumergroup"
)

// use it just like sync.Once
var once sync.Once

// ProcessEvent structure use to store the ProcessData fonction use to process the data conusme by the consumer
type ProcessEvent struct {
	ProcessData func(data interface{})
}

// eventMap is a map of event and ProcessEvent, to link
// the event comsume by the consumer and the function to process the data
type eventMap map[string]*ProcessEvent

// Process find the process data function by event name
func (objectMap *eventMap) Process(eventName string) (*ProcessEvent, error) {
	out := (*objectMap)[eventName]
	if out == nil {
		log.Warn("Can not found event named: '", eventName, "'")
		return nil, errors.New("Event not found : " + eventName)
	}
	return out, nil
}

// Consumer object that fetch the informations from kafka and store them in different channel.
// Consumer get the informations from a specific Topic
type Consumer struct {
	Topic        string
	Group        string
	BrokerList   []string
	Config       *consumergroup.Config
	Consumer     *consumergroup.ConsumerGroup
	GetModel     func(value []byte) interface{}
	GetEventType func(interface{}) string
	eMap         *eventMap
}

// RegisterConsumer register a new kaka consumer
func RegisterConsumer(registeredConsumer *Consumer) (*Consumer, error) {
	topics := []string{registeredConsumer.Topic}
	var consumerErr error
	registeredConsumer.Consumer, consumerErr = consumergroup.JoinConsumerGroup(
		registeredConsumer.Group,
		topics,
		registeredConsumer.BrokerList,
		registeredConsumer.Config)
	if consumerErr != nil {
		log.Fatalln(consumerErr)
		return nil, consumerErr
	}
	go func() {
		for {
			registeredConsumer.ProcessKafka()
		}
	}()
	return registeredConsumer, nil
}

// RegisterEvent register the event into the map
func (objectConsumer *Consumer) RegisterEvent(eventName string, object *ProcessEvent) {
	log.Info("Register event ", eventName, " for ", objectConsumer.Topic)
	(*objectConsumer.eventMap())[eventName] = object
}

// eventMap event map of the webhook
func (objectConsumer *Consumer) eventMap() *eventMap {
	if objectConsumer.eMap == nil {
		once.Do(func() {
			log.Debug("Alloc internal event map for consumer:", objectConsumer.Topic)
			eventMaping := make(eventMap)
			objectConsumer.eMap = &eventMaping
		})
	}
	return objectConsumer.eMap
}

// ProcessKafka consume th einformation from a topic kafka
func (objectConsumer *Consumer) ProcessKafka() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		if err := objectConsumer.Consumer.Close(); err != nil {
			sarama.Logger.Println("Error closing the consumer", err)
		}
	}()
	go func() {
		for err := range objectConsumer.Consumer.Errors() {
			log.Fatal(err)
		}
	}()
	for message := range objectConsumer.Consumer.Messages() {
		log.Print("Processing message with topic: ", message.Topic, " with partition: ", message.Partition, " and offset: ", message.Offset)
		model := objectConsumer.GetModel(message.Value)
		processEvent, err := objectConsumer.eventMap().Process(objectConsumer.GetEventType(model))
		if err != nil {
			log.Warn("unknow event found")
			continue
		}
		processEvent.ProcessData(model)
		time.Sleep(20 * time.Millisecond)
		objectConsumer.Consumer.CommitUpto(message)
	}
}
