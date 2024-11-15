package models

import (
    "github.com/gorilla/websocket"
    "sync"
)

type Stream struct {
    ID       string
    Conn     *websocket.Conn
    ConnLock sync.Mutex
}

var (
    Streams      = make(map[string]*Stream)
    StreamsMutex = &sync.Mutex{}
)

func (s *Stream) SetConnection(conn *websocket.Conn) {
    s.ConnLock.Lock()
    defer s.ConnLock.Unlock()
    s.Conn = conn
}