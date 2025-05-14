package main

import (
	"fmt"
	"net/http"
	"strings"
)

func fetchData(url string) (int, error) {
	// Make the HTTP GET request
	res, err := http.Get(url)
	if err != nil {
		// Handle network error
		return 0, fmt.Errorf("network error: %v", err)
	}
	defer res.Body.Close() // Ensure the response body is closed

	// Check if the status code is not 200
	if res.StatusCode != http.StatusOK {
		return res.StatusCode, fmt.Errorf("non-OK HTTP status: %s", res.Status)
	}

	// Return the status code and nil if no errors occurred
	return res.StatusCode, nil
}

func sendPostRequest() error {
	url := "https://api.boot.dev/v1/courses_rest_api/learn-http/users"
	payload := `{
		"role": "QA Job Safety",
		"experience": 2,
		"remote": true,
		"user": {
			"name": "Dan",
			"location": "NOR",
			"age": 29
		}
	}`

	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("non-OK HTTP status: %s", res.Status)
	}

	fmt.Println("Request successful")
	return nil
}
