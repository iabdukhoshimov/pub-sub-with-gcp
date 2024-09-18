package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"go.uber.org/zap"
)

type PubSubConfig struct {
	Client         *pubsub.Client
	Topic          *pubsub.Topic
	Subscription   *pubsub.Subscription
	SubscriptionID string
	TopicID        string
	Logger         *zap.Logger
}

// NewPubSubConfig initializes the Pub/Sub config with a logger
func NewPubSubConfig(client *pubsub.Client, topicID, subscriptionID string, logger *zap.Logger) *PubSubConfig {
	return &PubSubConfig{
		Client:         client,
		TopicID:        topicID,
		SubscriptionID: subscriptionID,
		Logger:         logger,
	}
}

func InitClient(ctx context.Context, projectID string) (*pubsub.Client, error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create pubsub client: %v", err)
	}
	return client, nil
}

// GetOrCreateTopic ensures that the topic exists or creates it
func (cfg *PubSubConfig) GetOrCreateTopic(ctx context.Context) error {
	topic := cfg.Client.Topic(cfg.TopicID)
	exists, err := topic.Exists(ctx)
	if err != nil {
		cfg.Logger.Error("Failed to check if topic exists", zap.Error(err))
		return err
	}
	if !exists {
		topic, err = cfg.Client.CreateTopic(ctx, cfg.TopicID)
		if err != nil {
			cfg.Logger.Error("Failed to create topic", zap.Error(err))
			return err
		}
		cfg.Logger.Info("Topic created", zap.String("topicID", cfg.TopicID))
	} else {
		cfg.Logger.Info("Topic already exists", zap.String("topicID", cfg.TopicID))
	}
	cfg.Topic = topic
	return nil
}
