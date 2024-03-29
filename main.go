package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	// Define your API URL
	apiUrl := os.Getenv("SERVER_URL")

	// Create a ticker that ticks every minute
	ticker := time.NewTicker(1 * time.Minute)

	for range ticker.C {
		// Ping the API URL
		if err := pingAPI(apiUrl); err != nil {
			fmt.Printf("Failed to ping API: %v\n", err)
		} else {
			fmt.Printf("API pinged successfully at %s\n", time.Now().Format(time.RFC3339))
		}
	}
}

func pingAPI(apiUrl string) error {
	// Send an HTTP GET request to the API URL
	_, err := http.Get(apiUrl)
	return err
}
