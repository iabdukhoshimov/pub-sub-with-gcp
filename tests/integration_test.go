package tests

import (
	"context"
	"testing"
	"time"

	"github.com/iabdukhoshimov/pubsub-microservice-golang/pkg/handlers"
	"github.com/iabdukhoshimov/pubsub-microservice-golang/pkg/pubsub"
	"github.com/stretchr/testify/assert"
)

func TestPubSubIntegration(t *testing.T) {
	ctx := context.Background()
	client, err := pubsub.InitClient(ctx, "test-project")
	assert.NoError(t, err)

	// Create or get the topic
	topic, err := pubsub.GetOrCreateTopic(ctx, client, "test-integration-topic")
	assert.NoError(t, err)

	// Publish a test message
	messageData := []byte(`{
		"id": "0001",
		"type": "donut",
		"name": "Cake",
		"image": {"url": "images/0001.jpg", "width": 200, "height": 200},
		"thumbnail": {"url": "images/thumbnails/0001.jpg", "width": 32, "height": 32}
	}`)
	err = pubsub.PublishMessage(ctx, topic, messageData)
	assert.NoError(t, err)

	// Create a subscription and start receiving messages
	done := make(chan bool)

	go func() {
		err := pubsub.SubscribeToTopic(ctx, client, "test-integration-subscription", topic, func(data []byte) {
			handlers.HandleMessage(data) // Handle message using the message handler
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
