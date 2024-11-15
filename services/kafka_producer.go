package services

import (
    "context"
    "encoding/json"
    "golang-redpanda-streaming/config"
	"log"
    "github.com/segmentio/kafka-go"
)

func PublishToKafka(streamID string, data map[string]interface{}) error {
    // Convert data to JSON
    message, err := json.Marshal(data)
    if err != nil {
        log.Printf("Error marshalling data: %v", err) // Add logging
        return err
    }

    // Create Kafka writer
    writer := kafka.NewWriter(kafka.WriterConfig{
        Brokers:  config.KafkaBrokers,
        Topic:    streamID,
        Balancer: &kafka.LeastBytes{},
    })
    defer writer.Close()

    // Write message to Kafka
    err = writer.WriteMessages(context.Background(),
        kafka.Message{
            Value: message,
        },
    )
    if err != nil {
        log.Printf("Error writing message to Kafka: %v", err) // Add logging
    }

    return err
}