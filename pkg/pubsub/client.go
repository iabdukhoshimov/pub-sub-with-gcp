package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

type PubSubConfig struct {
	Client         *pubsub.Client
	Topic          *pubsub.Topic
	Subscription   *pubsub.Subscription
	SubscriptionID string
	TopicID        string
}

// NewPubSubConfig initializes the configuration with the client, topic, and subscription IDs
func NewPubSubConfig(client *pubsub.Client, topicID string, subscriptionID string) *PubSubConfig {
	return &PubSubConfig{
		Client:         client,
		TopicID:        topicID,
		SubscriptionID: subscriptionID,
	}
}

// InitClient initializes the Pub/Sub client
func InitClient(ctx context.Context, projectID string) (*pubsub.Client, error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create pubsub client: %v", err)
	}
	return client, nil
}

// GetOrCreateTopic ensures that a topic exists or creates it
func (cfg *PubSubConfig) GetOrCreateTopic(ctx context.Context) error {
	topic := cfg.Client.Topic(cfg.TopicID)
	exists, err := topic.Exists(ctx)
	if err != nil {
		return fmt.Errorf("failed to check if topic exists: %v", err)
	}
	if !exists {
		topic, err = cfg.Client.CreateTopic(ctx, cfg.TopicID)
		if err != nil {
			return fmt.Errorf("failed to create topic: %v", err)
		}
		fmt.Printf("Topic %s created\n", cfg.TopicID)
	} else {
		fmt.Printf("Topic %s already exists\n", cfg.TopicID)
	}
	cfg.Topic = topic
	return nil
}
