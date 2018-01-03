package kafka

import (
	"os"
	"testing"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/smartystreets/goconvey/convey"
	"github.com/wvanbergen/kafka/consumergroup"
)

var (
	consumer           *Consumer
	consumerBrokerList []string
	configConsumer     *consumergroup.Config
	testObject         *testing.T
	topicsInit2        []string
)

func init() {
	zookeeperHost := os.Getenv("ZOOKEEPER_TEST")
	if zookeeperHost == "" {
		log.Print("Please set up ZOOKEEPER_TEST variable")
		// panic("No kafka host specify")
		return
	}
	kafkaHost := os.Getenv("KAFKA_TEST")
	if kafkaHost == "" {
		log.Print("Please set up KAFKA_TEST variable")
		// panic("No kafka host specify")
		return
	}
	brokerList = append(brokerList, kafkaHost)
	topicsInit2 = append(topicsInit2, "test2")
	producer.NewProducer(brokerList, topicsInit2, config)
	configConsumer = consumergroup.NewConfig()
	consumerBrokerList = append(consumerBrokerList, zookeeperHost)
	consumer, _ = RegisterConsumer(&Consumer{
		Topic:        "test2",
		Group:        "test2",
		BrokerList:   consumerBrokerList,
		Config:       configConsumer,
		GetModel:     testModel,
		GetEventType: testEventName,
	})
	consumer.RegisterEvent("testEvent", &ProcessEvent{testProcess})
}

func testModel(value []byte) interface{} {
	return string(value[:])
}

func testEventName(event interface{}) string {
	return "testEvent"
}

func testProcess(testEvent interface{}) {
	convey.Convey("testEvent should equal init message ", testObject, func() {
		convey.So(testEvent, convey.ShouldEqual, "init message")
	})
}

func TestConsumer(t *testing.T) {
	if consumer == nil {
		return
	}
	testObject = t
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				consumer.ProcessKafka()
			}
		}
	}()
	time.Sleep(2 * time.Second)
	close(quit)
}
