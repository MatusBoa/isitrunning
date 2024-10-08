package kafka

import (
	"fmt"
	"isitrunning/backend/events"

	"github.com/IBM/sarama"
)

// ====================
// Event Dispatcher
// ====================
type KafkaEventDispatcher struct {
	producer sarama.SyncProducer
}

func (dispatcher KafkaEventDispatcher) Dispatch(topic string, event events.Event) error {
	_, _, err := dispatcher.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(event.ToString()),
	})

	return err
}

// ====================
// Helper functions
// ====================
func CreateEventDispatcher(serverAddress string) KafkaEventDispatcher {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{serverAddress}, config)

	if err != nil {
		panic(fmt.Errorf("failed to setup producer: %w", err))
	}

	return KafkaEventDispatcher{
		producer: producer,
	}
}

