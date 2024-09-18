package main

import (
	"context"
	"log"

	"github.com/iabdukhoshimov/pubsub-microservice-golang/pkg/handlers"
	"github.com/iabdukhoshimov/pubsub-microservice-golang/pkg/pubsub"
)

const (
	projectID      = "your-gcp-project-id"
	topicID        = "example-topic"
	subscriptionID = "example-subscription"
)

func main() {
	ctx := context.Background()

	// Initialize Pub/Sub client
	client, err := pubsub.InitClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to initialize Pub/Sub client: %v", err)
	}
	defer client.Close()

	// Create Pub/Sub config
	cfg := pubsub.NewPubSubConfig(client, topicID, subscriptionID)

	// Create or get the topic
	err = cfg.GetOrCreateTopic(ctx)
	if err != nil {
		log.Fatalf("Failed to get or create topic: %v", err)
	}

	// Publish a message
	err = cfg.PublishMessage(ctx, []byte(`{
		"id": "0001",
		"type": "donut",
		"name": "Cake",
		"image": {"url": "images/0001.jpg", "width": 200, "height": 200},
		"thumbnail": {"url": "images/thumbnails/0001.jpg", "width": 32, "height": 32}
	}`))
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	// Subscribe and handle messages
	err = cfg.SubscribeToTopic(ctx, handlers.HandleMessage)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}
}
