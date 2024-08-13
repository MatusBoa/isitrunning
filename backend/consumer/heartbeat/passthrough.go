package heartbeat

import (
	"encoding/json"
	"isitrunning/backend/consumer"
	"isitrunning/backend/events"
	"isitrunning/backend/websockets/pusher"
	"log"

	"github.com/IBM/sarama"
)

type HearthbeatPassthroughConsumer struct {
	ConsumerConfig consumer.ConsumerConfig
}

func (*HearthbeatPassthroughConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (*HearthbeatPassthroughConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c HearthbeatPassthroughConsumer) Config() consumer.ConsumerConfig {
	return c.ConsumerConfig
}

func (consumer *HearthbeatPassthroughConsumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// @todo: Use env
	ws := pusher.CreateWebsocketClient("app-id", "app-secret", "app-key", "127.0.0.1:6001")

	for msg := range claim.Messages() {
		event := events.HeartbeatEvent{}

		err := json.Unmarshal(msg.Value, &event)

		if err != nil {
			log.Fatal(err)
			continue
		}

		ws.Emit(event.MonitorUuid, "heartbeat", string(msg.Value))
		log.Println("passthrough")
	}

	return nil
}
