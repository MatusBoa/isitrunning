package producer

import (
	"fmt"

	"github.com/IBM/sarama"
)

type Producer struct {
	client sarama.SyncProducer
}

func CreateProducer(kafkaServerAddress string) Producer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{kafkaServerAddress}, config)

	if err != nil {
		panic(fmt.Errorf("failed to setup producer: %w", err))
	}

	return Producer{
		client: client,
	}
}

func (p *Producer) Send(topic string, message string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err := p.client.SendMessage(msg)

	return err
}

func (p *Producer) Close() {
	p.client.Close()
}