package services

import (
    "context"
    "log"
    "golang-redpanda-streaming/config"
    "golang-redpanda-streaming/models"
    "github.com/segmentio/kafka-go"
)

func StartConsumer(streamID string, stream *models.Stream) {
    reader := kafka.NewReader(kafka.ReaderConfig{
        Brokers:  config.KafkaBrokers,
        GroupID:  streamID + "-group",
        Topic:    streamID,
        MinBytes: 1,
        MaxBytes: 10e6,
    })
    defer reader.Close()

    for {
        m, err := reader.ReadMessage(context.Background())
        if err != nil {
            log.Printf("Error reading message for stream %s: %v", streamID, err)
            return
        }

        // Process the message
        processedData, err := ProcessData(m.Value)
        if err != nil {
            log.Printf("Error processing data for stream %s: %v", streamID, err)
            continue
        }

        // Send processed data back to client
        stream.ConnLock.Lock()
        if stream.Conn != nil {
            err = stream.Conn.WriteJSON(processedData)
            if err != nil {
                log.Printf("Error sending data to client for stream %s: %v", streamID, err)
                stream.Conn = nil
            }
        }
        stream.ConnLock.Unlock()
    }
}

