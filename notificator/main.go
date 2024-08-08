package main

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/pusher/pusher-http-go/v5"
)

const (
	ConsumerGroup      = "page-group2"
	ConsumerTopic      = "page"
	KafkaServerAddress = "localhost:9092"
)

var (
	PusherClient pusher.Client
)

type Consumer struct {
}

func (*Consumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*Consumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (consumer *Consumer) ConsumeClaim(
	sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Print(string(msg.Value))
		PusherClient.Trigger("test", "stats", string(msg.Value))
	}
	return nil
}

func initializeConsumerGroup() (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()

	consumerGroup, err := sarama.NewConsumerGroup(
		[]string{KafkaServerAddress}, ConsumerGroup, config)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize consumer group: %w", err)
	}

	return consumerGroup, nil
}

func setupConsumerGroup(ctx context.Context) {
	consumerGroup, err := initializeConsumerGroup()
	if err != nil {
		log.Printf("initialization error: %v", err)
	}
	defer consumerGroup.Close()

	consumer := &Consumer{}

	for {
		err = consumerGroup.Consume(ctx, []string{ConsumerTopic}, consumer)
		if err != nil {
			log.Printf("error from consumer: %v", err)
		}
		if ctx.Err() != nil {
			return
		}
	}
}

func main() {
	PusherClient = pusher.Client{
		AppID:  "app-id",
		Secret: "app-secret",
		Key:    "app-key",
		Host:   "127.0.0.1:6001",
	}

	//data := map[string]string{
	//	"message": "hello world",
	//}

	//triggerErr := client.Trigger("my-channel", "my_event", data)

	//if triggerErr != nil {
	//panic(triggerErr)
	//}

	ctx, cancel := context.WithCancel(context.Background())
	setupConsumerGroup(ctx)
	defer cancel()
}
