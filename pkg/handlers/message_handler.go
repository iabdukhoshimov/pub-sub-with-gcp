package handlers

import (
	"encoding/json"

	"go.uber.org/zap"
)

// MessagePayload defines the structure of the incoming message
type MessagePayload struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Image     Image  `json:"image"`
	Thumbnail Image  `json:"thumbnail"`
}

// Image defines the structure of image fields in the message
type Image struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// HandleMessage is the function that processes a received message
func HandleMessage(logger *zap.Logger, data []byte) {
	// Parse the JSON message into the MessagePayload struct
	var payload MessagePayload
	err := json.Unmarshal(data, &payload)
	if err != nil {
		logger.Error("Error parsing message", zap.Error(err))
		return
	}

	// Extract and log the necessary fields using Zap
	logger.Info("Received message",
		zap.String("ID", payload.ID),
		zap.String("Type", payload.Type),
		zap.String("Image URL", payload.Image.URL),
		zap.String("Thumbnail URL", payload.Thumbnail.URL),
	)
}
