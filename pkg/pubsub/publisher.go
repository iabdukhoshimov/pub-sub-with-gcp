package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

// PublishMessage publishes a message to the topic
func (cfg *PubSubConfig) PublishMessage(ctx context.Context, data []byte) error {
	if cfg.Topic == nil {
		return fmt.Errorf("topic not initialized")
	}
	message := &pubsub.Message{
		Data: data,
	}

	result := cfg.Topic.Publish(ctx, message)
	_, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	fmt.Println("Message published successfully")
	return nil
}
