package controllers

import (
    "encoding/json"
    "net/http"
    "golang-redpanda-streaming/models"
    "golang-redpanda-streaming/services"
    "golang-redpanda-streaming/utils"

    "github.com/gorilla/mux"
    "github.com/gorilla/websocket"
)

func StartStream(w http.ResponseWriter, r *http.Request) {
    // Generate a unique stream ID
    streamID := utils.GenerateStreamID()

    stream := &models.Stream{
        ID: streamID,
    }

    models.StreamsMutex.Lock()
    models.Streams[streamID] = stream
    models.StreamsMutex.Unlock()

    go services.StartConsumer(streamID, stream)

    response := map[string]string{"stream_id": streamID}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func SendData(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    streamID := vars["stream_id"]

    var data map[string]interface{}
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, "Invalid data", http.StatusBadRequest)
        return
    }

    err = services.PublishToKafka(streamID, data)
    if err != nil {
        http.Error(w, "Failed to publish data", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func GetResults(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    streamID := vars["stream_id"]

    // Get the stream
    models.StreamsMutex.Lock()
    stream, exists := models.Streams[streamID]
    models.StreamsMutex.Unlock()

    if !exists {
        http.Error(w, "Stream not found", http.StatusNotFound)
        return
    }

    // Upgrade HTTP connection to WebSocket
    upgrader := websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool {
            return true
        },
    }
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
        return
    }

    // Set the WebSocket connection
    stream.SetConnection(conn)
}