package controllers

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"take-home-challenge/models"
)

func TestUsersController_Routes(t *testing.T) {
	user := &models.User{
		ID:        1,
		ClientID:  1,
		FirstName: "John",
		LastName:  "Doe",
		Metadata: map[string]interface{}{
			"tag": "employee",
		},
	}
	url := fmt.Sprintf("%s/api/v2/users/", server.URL)
	req, _ := http.NewRequest(http.MethodPost, url, toBody(user))
	resp, _ := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	url = fmt.Sprintf("%s/api/v2/users/1", server.URL)
	req, _ = http.NewRequest(http.MethodGet, url, nil)
	resp, _ = http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
	u := fromBody[*models.User](resp)
	if !reflect.DeepEqual(u, user) {
		t.Fatalf("Expected user %v, got %v", user, u)
	}
}
