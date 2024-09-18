package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"go.uber.org/zap"
)

func (cfg *PubSubConfig) PublishMessage(ctx context.Context, data []byte) error {
	if cfg.Topic == nil {
		cfg.Logger.Error("Topic not initialized")
		return fmt.Errorf("topic not initialized")
	}
	message := &pubsub.Message{
		Data: data,
	}

	result := cfg.Topic.Publish(ctx, message)
	_, err := result.Get(ctx)
	if err != nil {
		cfg.Logger.Error("Failed to publish message", zap.Error(err))
		return err
	}

	cfg.Logger.Info("Message published successfully")
	return nil
}
