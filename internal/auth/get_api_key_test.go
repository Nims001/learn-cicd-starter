package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	Apikey := "AppKey crazyapi6969" // Set the API key for testing

	// 1. Create a new HTTP request
	req, err := http.NewRequest("GET", "https://api.example.com/data", nil)
	if err != nil {
		panic(err)
	}

	// 2. Set or add your headers
	req.Header.Set("Authorization", Apikey)      // Overwrites existing keys
	req.Header.Add("Accept", "application/json") // Appends to existing keys
	req.Header.Set("User-Agent", "Go-Client/1.0")

	res, err := GetAPIKey(req.Header)
	if err != nil {
		t.Fatalf("Failed to get API key: %v", err)
	}

	// 4. Check the response
	if res != Apikey {
		t.Errorf("Expected status OK, got %v", res)
	}
}
