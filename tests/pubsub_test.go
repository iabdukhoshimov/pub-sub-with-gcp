package tests

import (
	"context"
	"testing"

	pubsub "github.com/iabdukhoshimov/pubsub-microservice-golang/pkg/pubsub"
	"github.com/stretchr/testify/assert"
)

func TestInitClient(t *testing.T) {
	ctx := context.Background()

	// Test successful client creation
	client, err := pubsub.InitClient(ctx, "test-project")
	assert.NoError(t, err, "Error should be nil when initializing Pub/Sub client")
	assert.NotNil(t, client, "Client should not be nil")
}

func TestGetOrCreateTopic(t *testing.T) {
	ctx := context.Background()
	client, err := pubsub.InitClient(ctx, "test-project")
	assert.NoError(t, err)

	// Create the PubSubConfig struct
	cfg := pubsub.NewPubSubConfig(client, "test-topic", "test-subscription")

	// Test topic creation
	err = cfg.GetOrCreateTopic(ctx)
	assert.NoError(t, err, "Error should be nil when creating topic")
	assert.NotNil(t, cfg.Topic, "Topic should not be nil")
}

func TestPublishMessage(t *testing.T) {
	ctx := context.Background()
	client, err := pubsub.InitClient(ctx, "test-project")
	assert.NoError(t, err)

	// Create PubSubConfig struct
	cfg := pubsub.NewPubSubConfig(client, "test-topic", "test-subscription")

	// Ensure the topic is created
	err = cfg.GetOrCreateTopic(ctx)
	assert.NoError(t, err)

	// Test publishing a message
	messageData := []byte(`{"id": "test-id", "type": "test-type"}`)
	err = cfg.PublishMessage(ctx, messageData)
	assert.NoError(t, err, "Error should be nil when publishing message")
}
