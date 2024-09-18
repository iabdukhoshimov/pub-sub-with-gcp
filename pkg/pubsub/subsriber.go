package pubsub

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
)

// SubscribeToTopic subscribes to a Pub/Sub topic and processes messages
func (cfg *PubSubConfig) SubscribeToTopic(ctx context.Context, messageHandler func([]byte)) error {
	sub := cfg.Client.Subscription(cfg.SubscriptionID)
	exists, err := sub.Exists(ctx)
	if err != nil {
		return fmt.Errorf("failed to check subscription existence: %v", err)
	}
	if !exists {
		sub, err = cfg.Client.CreateSubscription(ctx, cfg.SubscriptionID, pubsub.SubscriptionConfig{
			Topic:       cfg.Topic,
			AckDeadline: 10 * time.Second,
		})
		if err != nil {
			return fmt.Errorf("failed to create subscription: %v", err)
		}
		fmt.Printf("Subscription %s created\n", cfg.SubscriptionID)
	}
	cfg.Subscription = sub

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		messageHandler(msg.Data)
		msg.Ack() // Acknowledge the message
	})
	if err != nil {
		return fmt.Errorf("failed to receive messages: %v", err)
	}

	return nil
}
