package data

import (
	"database/sql"
	"errors"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrConcurrentWrite = errors.New("concurrent write")
)

type DB interface {
	Reader() *sql.DB
	Writer() *sql.DB
	Ping() error
}

type AggregateEvent struct {
	ID          string `json:"id"`
	AggregateID string `json:"aggregate_id"`
	Timestamp   int64  `json:"timestamp"`
	Version     uint32 `json:"version"`
	EventType   string `json:"type"`
	Event       []byte `json:"event"`
}

type StreamResult struct {
	AggregateEvent AggregateEvent
	Err            error
}

type Tx interface {
	Rollback() error
	Commit() error
}

type Repository interface {
	Read(id string) ([]AggregateEvent, error)
	Write(id string, events []AggregateEvent, version uint32) (Tx, error)
	Stream(id string, lastSeen uint32, ch chan<- StreamResult)
}
