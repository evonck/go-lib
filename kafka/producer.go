package kafka

import (
	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
)

// Producer object comunicate with kafka consumer
type Producer struct {
	producer sarama.SyncProducer
}

// Close the producer
func (k *Producer) Close() {
	if err := k.producer.Close(); err != nil {
		log.Println("Failed to shut down data collector cleanly", err)
	}
}

// InitKafka initialise a list of topics in kafka
func (k *Producer) InitKafka(topics []string) {
	for _, topic := range topics {
		_, _, err := k.producer.SendMessage(&sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder("init message"),
		})
		if err != nil {
			log.Warnf("Failed to send your data %s", err)
			panic("error during init")
		}
		log.Print("init topic : " + topic)
	}
}

// NewProducer create a new kafka producer
func (k *Producer) NewProducer(brokerList, topics []string, config *sarama.Config) error {
	var err error
	k.producer, err = sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Error("Failed to start producer:", err)
		return err
	}
	k.InitKafka(topics)
	return nil
}

// SendData send message to kafka topic
func (k *Producer) SendData(topic, message string) {
	partition, offset, err := k.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	})
	if err != nil {
		log.Warnf("Failed to send your data %s", err)
	} else {
		log.Warnf("Your data is stored with unique topic %s/ partition %d/ offset %d", topic, partition, offset)
	}
}
