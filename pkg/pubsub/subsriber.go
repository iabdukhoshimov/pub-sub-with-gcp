package pubsub

import (
	"context"
	"time"

	"cloud.google.com/go/pubsub"
	"go.uber.org/zap"
)

// SubscribeToTopic subscribes to a Pub/Sub topic and processes messages
func (cfg *PubSubConfig) SubscribeToTopic(ctx context.Context, messageHandler func([]byte)) error {
	sub := cfg.Client.Subscription(cfg.SubscriptionID)
	exists, err := sub.Exists(ctx)
	if err != nil {
		cfg.Logger.Error("Failed to check subscription existence", zap.Error(err))
		return err
	}
	if !exists {
		sub, err = cfg.Client.CreateSubscription(ctx, cfg.SubscriptionID, pubsub.SubscriptionConfig{
			Topic:       cfg.Topic,
			AckDeadline: 10 * time.Second,
		})
		if err != nil {
			cfg.Logger.Error("Failed to create subscription", zap.Error(err))
			return err
		}
		cfg.Logger.Info("Subscription created", zap.String("subscriptionID", cfg.SubscriptionID))
	}
	cfg.Subscription = sub

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		cfg.Logger.Info("Message received", zap.String("messageID", msg.ID))
		messageHandler(msg.Data)
		msg.Ack() // Acknowledge the message
	})
	if err != nil {
		cfg.Logger.Error("Failed to receive messages", zap.Error(err))
		return err
	}

	return nil
}
