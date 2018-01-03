package kafka

import (
	"os"
	"testing"

	"bitbucket.org/onekloud/go-lib/tools"
	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
	"github.com/smartystreets/goconvey/convey"
)

var (
	producer   Producer
	brokerList []string
	topicsInit []string
	config     *sarama.Config
)

func init() {
	kafkaHost := os.Getenv("KAFKA_TEST")
	if kafkaHost == "" {
		log.Print("Please set up KAFKA_TEST variable")
		return
		// panic("No kafka host specify")
	}
	brokerList = append(brokerList, kafkaHost)
	topicsInit = append(topicsInit, "test")
}

func TestInit(t *testing.T) {
	err := producer.NewProducer(brokerList, topicsInit, config)
	if brokerList == nil {
		convey.Convey("err should not be nil", t, func() {
			convey.So(err, convey.ShouldNotEqual, nil)
		})
		return
	}
	kafkaClient, err := sarama.NewClient(brokerList, config)
	if err != nil {
		panic(err)
	}
	defer kafkaClient.Close()
	topics, err := kafkaClient.Topics()
	if err != nil {
		panic(err)
	}
	for _, topic := range topics {
		convey.Convey("config topic should be in kafka topics ", t, func() {
			convey.So(tools.IsInArray(topic, topics), convey.ShouldEqual, true)
		})
	}
}

func TestSendData(t *testing.T) {
	kafkaClient, err := sarama.NewClient(brokerList, config)
	if err != nil {
		log.Info(err)
		return
	}
	defer kafkaClient.Close()
	partitionID, err := kafkaClient.Partitions(topicsInit[0])
	convey.Convey("err should be nil", t, func() {
		convey.So(err, convey.ShouldEqual, nil)
	})
	convey.Convey("partitionID should not be nil ", t, func() {
		convey.So(partitionID, convey.ShouldNotEqual, nil)
	})
	offset, err := kafkaClient.GetOffset("test", partitionID[0], sarama.OffsetOldest)
	convey.Convey("err should be nil", t, func() {
		convey.So(err, convey.ShouldEqual, nil)
	})
	producer.NewProducer(brokerList, topicsInit, config)
	producer.SendData(topicsInit[0], "init message")
	offset2, err := kafkaClient.GetOffset("test", partitionID[0], sarama.OffsetOldest)
	convey.Convey("err should be nil", t, func() {
		convey.So(err, convey.ShouldEqual, nil)
	})
	if offset == 0 {
		convey.Convey("offset2 should not be equal to offset ", t, func() {
			convey.So(offset2, convey.ShouldEqual, offset)
		})
	} else {
		convey.Convey("offset2 should not be equal to offset + 1 ", t, func() {
			convey.So(offset2, convey.ShouldEqual, offset+1)
		})
	}
}
