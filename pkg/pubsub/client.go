package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

// InitClient initializes the Pub/Sub client
func InitClient(ctx context.Context, projectID string) (*pubsub.Client, error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create pubsub client: %v", err)
	}
	return client, nil
}

// GetOrCreateTopic checks if the topic exists and creates it if not
func GetOrCreateTopic(ctx context.Context, client *pubsub.Client, topicID string) (*pubsub.Topic, error) {
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
