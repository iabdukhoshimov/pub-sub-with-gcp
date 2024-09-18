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

	// Test topic creation
	topic, err := pubsub.GetOrCreateTopic(ctx, client, "test-topic")
	assert.NoError(t, err, "Error should be nil when creating topic")
	assert.NotNil(t, topic, "Topic should not be nil")
}

func TestPublishMessage(t *testing.T) {
	ctx := context.Background()
	client, err := pubsub.InitClient(ctx, "test-project")
	assert.NoError(t, err)

	topic, err := pubsub.GetOrCreateTopic(ctx, client, "test-topic")
	assert.NoError(t, err)

	// Test publishing a message
	messageData := []byte(`{"id": "test-id", "type": "test-type"}`)
	err = pubsub.PublishMessage(ctx, topic, messageData)
	assert.NoError(t, err, "Error should be nil when publishing message")
}
