package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
)

const (
	projectID      = "your-gcp-project-id"  // Replace with your GCP project ID
	topicID        = "example-topic"        // Topic name
	subscriptionID = "example-subscription" // Subscription name
)

func main() {
	// Set up Pub/Sub client
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub client: %v", err)
	}
	defer client.Close()

	// Create a Pub/Sub topic
	topic, err := getOrCreateTopic(ctx, client, topicID)
	if err != nil {
		log.Fatalf("Failed to get or create topic: %v", err)
	}

	// Publish a message to the topic
	err = publishMessage(ctx, topic)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	// Subscribe to the topic and consume messages
	err = subscribeToTopic(ctx, client, subscriptionID, topic)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}
}

// getOrCreateTopic checks if a topic exists, and creates it if not
func getOrCreateTopic(ctx context.Context, client *pubsub.Client, topicID string) (*pubsub.Topic, error) {
	topic := client.Topic(topicID)
	exists, err := topic.Exists(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check if topic exists: %v", err)
	}
	if !exists {
		topic, err = client.CreateTopic(ctx, topicID)
		if err != nil {
			return nil, fmt.Errorf("failed to create topic: %v", err)
		}
		fmt.Printf("Topic %s created\n", topicID)
	} else {
		fmt.Printf("Topic %s already exists\n", topicID)
	}
	return topic, nil
}

// publishMessage publishes a message to a given topic
func publishMessage(ctx context.Context, topic *pubsub.Topic) error {
	message := &pubsub.Message{
		Data: []byte(`{
			"id": "0001",
			"type": "donut",
			"name": "Cake",
			"image": {"url": "images/0001.jpg", "width": 200, "height": 200},
			"thumbnail": {"url": "images/thumbnails/0001.jpg", "width": 32, "height": 32}
		}`),
	}

	result := topic.Publish(ctx, message)
	_, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}
	fmt.Println("Message published successfully")
	return nil
}

// subscribeToTopic subscribes to a topic and processes incoming messages
func subscribeToTopic(ctx context.Context, client *pubsub.Client, subscriptionID string, topic *pubsub.Topic) error {
	// Create a subscription if it doesn't exist
	sub := client.Subscription(subscriptionID)
	exists, err := sub.Exists(ctx)
	if err != nil {
		return fmt.Errorf("failed to check subscription existence: %v", err)
	}
	if !exists {
		sub, err = client.CreateSubscription(ctx, subscriptionID, pubsub.SubscriptionConfig{
			Topic:       topic,
			AckDeadline: 10 * time.Second,
		})
		if err != nil {
			return fmt.Errorf("failed to create subscription: %v", err)
		}
		fmt.Printf("Subscription %s created\n", subscriptionID)
	} else {
		fmt.Printf("Subscription %s already exists\n", subscriptionID)
	}

	// Consume messages from the subscription
	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("Received message: %s\n", string(msg.Data))

		// Here, parse the message and extract necessary fields
		// In this case, we're just printing the message
		msg.Ack() // Acknowledge the message
	})
	if err != nil {
		return fmt.Errorf("failed to receive messages: %v", err)
	}

	return nil
}
