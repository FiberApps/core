package kafka

import (
	"fmt"

	"github.com/FiberApps/core/logger"
	"github.com/Shopify/sarama"
)

// Consumer
func createConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

// Producer
func createProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

// Publisher
func PublishMessage(kafkaBroker string, topic string, message []byte) error {
	log := logger.New()
	brokersUrl := []string{kafkaBroker}
	producer, err := createProducer(brokersUrl)
	if err != nil {
		return err
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}
	log.Info(fmt.Sprintf("Message is stored in topic(%s)/partition(%d)/offset(%d)", topic, partition, offset))
	return nil
}
