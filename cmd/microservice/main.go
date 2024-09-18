package main

import (
	"context"
	"log"

	"github.com/iabdukhoshimov/pubsub-microservice-golang/pkg/handlers"
	pubsub "github.com/iabdukhoshimov/pubsub-microservice-golang/pkg/pubsub"
	"go.uber.org/zap"
)

const (
	projectID      = "your-gcp-project-id"
	topicID        = "example-topic"
	subscriptionID = "example-subscription"
)

func main() {
	// Initialize Zap logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	ctx := context.Background()

	// Initialize Pub/Sub client
	client, err := pubsub.InitClient(ctx, projectID)
	if err != nil {
		logger.Fatal("Failed to initialize Pub/Sub client", zap.Error(err))
	}
	defer client.Close()

	// Create Pub/Sub config with logger
	cfg := pubsub.NewPubSubConfig(client, topicID, subscriptionID, logger)

	// Create or get the topic
	if err = cfg.GetOrCreateTopic(ctx); err != nil {
		logger.Fatal("Failed to get or create topic", zap.Error(err))
	}

	// Publish a message
	message := []byte(`{
		"id": "0001",
		"type": "donut",
		"name": "Cake",
		"image": {"url": "images/0001.jpg", "width": 200, "height": 200},
		"thumbnail": {"url": "images/thumbnails/0001.jpg", "width": 32, "height": 32}
	}`)
	if err = cfg.PublishMessage(ctx, message); err != nil {
		logger.Fatal("Failed to publish message", zap.Error(err))
	}

	// Subscribe to the topic and process messages
	if err = cfg.SubscribeToTopic(ctx, func(data []byte) {
		handlers.HandleMessage(logger, data) // Pass logger to HandleMessage
	}); err != nil {
		logger.Fatal("Failed to subscribe to topic", zap.Error(err))
	}
}
