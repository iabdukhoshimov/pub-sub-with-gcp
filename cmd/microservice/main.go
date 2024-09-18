package main

import (
	"context"
	"log"

	"github.com/iabdukhoshimov/pubsub-microservice-golang/pkg/handlers"
	"github.com/iabdukhoshimov/pubsub-microservice-golang/pkg/pubsub"
)

const (
	projectID      = "sample"
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

	// Create or get the topic
	topic, err := pubsub.GetOrCreateTopic(ctx, client, topicID)
	if err != nil {
		log.Fatalf("Failed to get or create topic: %v", err)
	}

	// Publish a message to the topic
	err = pubsub.PublishMessage(ctx, topic, []byte(`{
		"id": "0001",
		"type": "donut",
		"name": "Cake",
		"image": {"url": "images/0001.jpg", "width": 200, "height": 200},
		"thumbnail": {"url": "images/thumbnails/0001.jpg", "width": 32, "height": 32}
	}`))
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	// Subscribe to the topic and process messages using the handler
	err = pubsub.SubscribeToTopic(ctx, client, subscriptionID, topic, handlers.HandleMessage)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}
}
