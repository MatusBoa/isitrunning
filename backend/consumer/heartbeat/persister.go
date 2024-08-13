package heartbeat

import (
	"encoding/json"
	"isitrunning/backend/consumer"
	"isitrunning/backend/events"
	"log"

	"github.com/IBM/sarama"
)

type HeartBeatPersisterConsumer struct {
	ConsumerConfig consumer.ConsumerConfig
}

func (*HeartBeatPersisterConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (*HeartBeatPersisterConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c HeartBeatPersisterConsumer) Config() consumer.ConsumerConfig {
	return c.ConsumerConfig
}

func (consumer *HeartBeatPersisterConsumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		event := events.HeartbeatEvent{}

		err := json.Unmarshal(msg.Value, &event)

		if err != nil {
			log.Fatal(err)
			continue
		}
	}

	return nil
}