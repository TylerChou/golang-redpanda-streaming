package utils

import (
    "crypto/rand"
    "encoding/hex"
)

func GenerateStreamID() string {
    bytes := make([]byte, 16)
    rand.Read(bytes)
    return hex.EncodeToString(bytes)
}