package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

// PublishMessage publishes a message to the specified Pub/Sub topic
func PublishMessage(ctx context.Context, topic *pubsub.Topic, data []byte) error {
	message := &pubsub.Message{
		Data: data,
	}

	result := topic.Publish(ctx, message)
	_, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	fmt.Println("Message published successfully")
	return nil
}
