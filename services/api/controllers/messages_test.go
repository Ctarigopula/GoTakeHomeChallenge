package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"take-home-challenge/models"
	"take-home-challenge/services/api/controllers/payloads"
)

func TestMessagesController_Routes(t *testing.T) {
	baseURL := fmt.Sprintf("%s/api/v2/messages", server.URL)

	messages := []models.Message{
		{
			ID:        1,
			Created:   nil,
			DeletedAt: nil,
			UserIDs:   []int{1, 2, 3},
			Metadata:  map[string]interface{}{"tag": "employee"},
		},
		{
			ID:        2,
			Created:   nil,
			DeletedAt: nil,
			UserIDs:   []int{1, 2, 3},
			Metadata:  map[string]interface{}{"tag": "employee"},
		},
		{
			ID:        3,
			Created:   nil,
			DeletedAt: nil,
			UserIDs:   []int{1, 2, 3},
			Metadata:  map[string]interface{}{"tag": "employee"},
		},
	}

	// Create messages
	for _, msg := range messages {
		req, err := http.NewRequest(http.MethodPost, baseURL+"/", toBody(&msg))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			t.Fatalf("Expected status %d, got %d", http.StatusCreated, resp.StatusCode)
		}
	}

	// Delete messages 2 and 3
	deleteURL := fmt.Sprintf("%s/delete", baseURL)
	deletePayload := &payloads.MessagesMarkDeleted{
		IDs:         []int{2, 3},
		DeletedWhen: "2025-06-01",
	}

	req, err := http.NewRequest(http.MethodPost, deleteURL, toBody(deletePayload))
	if err != nil {
		t.Fatalf("Failed to create delete request: %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Delete request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Fetch message 1 and verify
	getURL := fmt.Sprintf("%s/1", baseURL)
	req, err = http.NewRequest(http.MethodGet, getURL, nil)
	if err != nil {
		t.Fatalf("Failed to create GET request: %v", err)
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("GET request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}
	// to check nil values
	var fetchedMsg models.Message
	if err := json.NewDecoder(resp.Body).Decode(&fetchedMsg); err != nil {
		t.Fatalf("Failed to decode message response: %v", err)
	}

	if fetchedMsg.Created != nil {
		t.Errorf("Expected Created to be nil, got %v", *fetchedMsg.Created)
	}
	if fetchedMsg.DeletedAt != nil {
		t.Errorf("Expected DeletedAt to be nil, got %v", *fetchedMsg.DeletedAt)
	}

	// Get deleted message (ID 2) - should return 404
	getURL = fmt.Sprintf("%s/2", baseURL)
	req, err = http.NewRequest(http.MethodGet, getURL, nil)
	if err != nil {
		t.Fatalf("Failed to create GET request: %v", err)
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("GET request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Expected 404 for deleted message, got %d", resp.StatusCode)
	}

	// Get invalid message (ID invalid) - should return 400
	getURL = fmt.Sprintf("%s/invalid", baseURL)
	req, err = http.NewRequest(http.MethodGet, getURL, nil)
	if err != nil {
		t.Fatalf("Failed to create GET request: %v", err)
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("GET request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("Expected 400 for invalid message id, got %d", resp.StatusCode)
	}

}
