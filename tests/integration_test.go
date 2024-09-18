package tests

import (
	"context"
	"testing"
	"time"

	"github.com/iabdukhoshimov/pubsub-microservice-golang/pkg/handlers"
	pubsub "github.com/iabdukhoshimov/pubsub-microservice-golang/pkg/pubsub"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestPubSubIntegration(t *testing.T) {
	ctx := context.Background()
	client, err := pubsub.InitClient(ctx, "test-project")
	assert.NoError(t, err)

	// Create test logger
	logger := zap.NewNop() // Nop logger for testing

	// Create PubSubConfig struct
	cfg := pubsub.NewPubSubConfig(client, "test-integration-topic", "test-integration-subscription", logger)

	// Create or get the topic
	err = cfg.GetOrCreateTopic(ctx)
	assert.NoError(t, err)

	// Publish a test message
	messageData := []byte(`{
		"id": "0001",
		"type": "donut",
		"name": "Cake",
		"image": {"url": "images/0001.jpg", "width": 200, "height": 200},
		"thumbnail": {"url": "images/thumbnails/0001.jpg", "width": 32, "height": 32}
	}`)
	err = cfg.PublishMessage(ctx, messageData)
	assert.NoError(t, err)

	// Create a subscription and start receiving messages
	done := make(chan bool)

	go func() {
		err := cfg.SubscribeToTopic(ctx, func(data []byte) {
			handlers.HandleMessage(logger, data) // Pass logger to HandleMessage
			done <- true
		})
		assert.NoError(t, err)
	}()

	// Wait for the message to be processed
	select {
	case <-done:
		// Test passes when the message is received and processed
	case <-time.After(10 * time.Second):
		t.Fatal("Test timed out, message not received")
	}
}
