package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    // Define your API URL
    apiUrl := "https://hng-ix-02.onrender.com"

    // Create a ticker that ticks every 2 minutes
    ticker := time.NewTicker(2 * time.Minute)

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
