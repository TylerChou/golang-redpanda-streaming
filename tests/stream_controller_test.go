package tests

import (
    "net/http"
    "net/http/httptest"
    "golang-redpanda-streaming/controllers"
    "testing"
)

func TestStartStream(t *testing.T) {
    req, err := http.NewRequest("POST", "/stream/start", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(controllers.StartStream)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
}