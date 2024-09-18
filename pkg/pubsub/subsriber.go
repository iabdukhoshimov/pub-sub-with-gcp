package pubsub

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
)

// SubscribeToTopic subscribes to a Pub/Sub topic and handles incoming messages
func SubscribeToTopic(ctx context.Context, client *pubsub.Client, subscriptionID string, topic *pubsub.Topic, messageHandler func([]byte)) error {
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

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		// Call the provided message handler function
		messageHandler(msg.Data)
		msg.Ack() // Acknowledge the message
	})
	if err != nil {
		return fmt.Errorf("failed to receive messages: %v", err)
	}

	return nil
}
