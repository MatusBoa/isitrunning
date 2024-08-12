package consumer

import (
	"log"

	"github.com/IBM/sarama"
)

type HeartbeatConsumer struct {
	ConsumerConfig ConsumerConfig
}

func (*HeartbeatConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (*HeartbeatConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c HeartbeatConsumer) Config() ConsumerConfig {
	return c.ConsumerConfig
}

func (consumer *HeartbeatConsumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Println(string(msg.Value))
	}

	return nil
}
