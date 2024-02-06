package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// read environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Define your API URL
	apiUrl := os.Getenv("URL")

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
