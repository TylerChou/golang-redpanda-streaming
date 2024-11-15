package routes

import (
    "github.com/gorilla/mux"
    "golang-redpanda-streaming/controllers"
    "golang-redpanda-streaming/middleware"
)

func InitRoutes() *mux.Router {
    router := mux.NewRouter()

    // Apply middleware
    router.Use(middleware.LoggingMiddleware)
    router.Use(middleware.AuthenticationMiddleware)

    // Define routes
    router.HandleFunc("/stream/start", controllers.StartStream).Methods("POST")
    router.HandleFunc("/stream/{stream_id}/send", controllers.SendData).Methods("POST")
    router.HandleFunc("/stream/{stream_id}/results", controllers.GetResults).Methods("GET")

    return router
}