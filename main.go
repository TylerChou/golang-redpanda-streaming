package main

import (
    "log"
    "net/http"
    "golang-redpanda-streaming/config"
    "golang-redpanda-streaming/routes"
)

func main() {
    // Initialize configuration
    config.LoadConfig()

    // Initialize routes
    router := routes.InitRoutes()

    // Start the server
    log.Println("Server is starting on port 8080...")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}