package handlers

import (
	"encoding/json"
	"fmt"
	"log"
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
func HandleMessage(data []byte) {
	// Parse the JSON message into the MessagePayload struct
	var payload MessagePayload
	err := json.Unmarshal(data, &payload)
	if err != nil {
		log.Printf("Error parsing message: %v", err)
		return
	}

	// Extract and print the necessary fields
	fmt.Printf("Received message:\n")
	fmt.Printf("ID: %s\n", payload.ID)
	fmt.Printf("Type: %s\n", payload.Type)
	fmt.Printf("Image URL: %s\n", payload.Image.URL)
	fmt.Printf("Thumbnail URL: %s\n", payload.Thumbnail.URL)
}
