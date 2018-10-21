package utils

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

func Timestamp() int64 {
	return time.Now().UTC().Unix()
}

func NewUUID() string {
	return uuid.NewV4().String()
}
