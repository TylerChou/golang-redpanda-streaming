package config

import (
    "os"
    "strings"
)

var (
    KafkaBrokers []string
    APIKey       string
)

func LoadConfig() {
    brokers := os.Getenv("KAFKA_BROKERS")
    if brokers == "" {
        brokers = "localhost:9092"
    }
    KafkaBrokers = strings.Split(brokers, ",")

    APIKey = os.Getenv("API_KEY")
    if APIKey == "" {
        APIKey = "default-api-key"
    }
}