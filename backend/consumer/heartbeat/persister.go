package heartbeat

import (
	"encoding/json"
	"isitrunning/backend/consumer"
	"isitrunning/backend/db"
	"isitrunning/backend/events"
	"isitrunning/backend/models"
	"isitrunning/backend/repositories"
	"isitrunning/backend/websockets/pusher"
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
	d, err := db.Initialize()

	if err != nil {
		panic(err)
	}

	mhr := repositories.CreateMonitorHeartbeatRepository(&d)
	ws := pusher.CreateWebsocketClient("app-id", "app-secret", "app-key", "127.0.0.1:6001")

	for msg := range claim.Messages() {
		event := events.HeartbeatEvent{}

		err := json.Unmarshal(msg.Value, &event)

		if err != nil {
			log.Fatal(err)
			continue
		}

		model := models.MonitorHeartbeat{
			MonitorUuid:  event.MonitorUuid,
			StatusCode:   event.StatusCode,
			ResponseTime: event.ResponseTime,
		}

		model = mhr.Insert(model)

		json, err := json.Marshal(model)

		if err != nil {
			panic(err)
		}

		ws.Emit(event.MonitorUuid, "heartbeat", string(json))
	}

	return nil
}
