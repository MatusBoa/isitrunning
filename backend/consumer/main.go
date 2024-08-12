package consumer

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
)

// ====================
// Types
// ====================
type ConsumerConfig struct {
	ServerAddress []string
	Group         string
	Topic         []string
}

type ConsumerCallback func(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error

// ====================
// Consumer
// ====================
type Consumer interface {
	sarama.ConsumerGroupHandler
	Config() ConsumerConfig
}

// ====================
// Internal functions
// ====================
func createConsumerGroup(constumerConfig ConsumerConfig) (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()

	consumerGroup, err := sarama.NewConsumerGroup(
		constumerConfig.ServerAddress,
		constumerConfig.Group,
		config,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize consumer group: %w", err)
	}

	return consumerGroup, nil
}

// ====================
// Helper functions
// ====================
func InitializeConsumer(
	context context.Context,
	consumer Consumer,
) {
	consumerGroup, err := createConsumerGroup(consumer.Config())

	if err != nil {
		panic(fmt.Errorf("initialization error: %v", err))
	}

	defer consumerGroup.Close()

	for {
		err = consumerGroup.Consume(context, consumer.Config().Topic, consumer)

		if err != nil {
			panic(fmt.Errorf("error from consumer: %v", err))
		}

		if context.Err() != nil {
			return
		}
	}
}
