package controllers

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"take-home-challenge/models"
	"take-home-challenge/services/api/controllers/payloads"
)

func TestMessagesController_Routes(t *testing.T) {
	message := &models.Message{
		ID:        1,
		Created:   nil,
		DeletedAt: nil,
		UserIDs:   []int{1, 2, 3},
		Metadata: map[string]interface{}{
			"tag": "employee",
		},
	}
	url := fmt.Sprintf("%s/api/v2/messages/", server.URL)
	req, _ := http.NewRequest(http.MethodPost, url, toBody(message))
	resp, _ := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	message = &models.Message{
		ID:        2,
		Created:   nil,
		DeletedAt: nil,
		UserIDs:   []int{1, 2, 3},
		Metadata: map[string]interface{}{
			"tag": "employee",
		},
	}
	url = fmt.Sprintf("%s/api/v2/messages/", server.URL)
	req, _ = http.NewRequest(http.MethodPost, url, toBody(message))
	resp, _ = http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	message = &models.Message{
		ID:        3,
		Created:   nil,
		DeletedAt: nil,
		UserIDs:   []int{1, 2, 3},
		Metadata: map[string]interface{}{
			"tag": "employee",
		},
	}
	url = fmt.Sprintf("%s/api/v2/messages/", server.URL)
	req, _ = http.NewRequest(http.MethodPost, url, toBody(message))
	resp, _ = http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	url = fmt.Sprintf("%s/api/v2/messages/3", server.URL)
	req, _ = http.NewRequest(http.MethodGet, url, nil)
	resp, _ = http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
	m := fromBody[*models.Message](resp)
	if !reflect.DeepEqual(m, message) {
		t.Fatalf("Expected user %v, got %v", message, m)
	}

	url = fmt.Sprintf("%s/api/v2/messages/delete", server.URL)
	data := &payloads.MessagesMarkDeleted{
		IDs:         []int{3, 2},
		DeletedWhen: "2025-06-01",
	}
	req, _ = http.NewRequest(http.MethodPost, url, toBody(data))
	resp, _ = http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

}
