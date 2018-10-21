package data

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type eventRepo struct {
	db         DB
	readStmt   *sql.Stmt
	streamStmt *sql.Stmt
	writeStmt  string
}

func NewEventRepository(db DB, table string) (*eventRepo, error) {
	r := &eventRepo{
		db:        db,
		writeStmt: fmt.Sprintf(writeStatement, table),
	}
	var err error
	if r.readStmt, err = r.db.Reader().Prepare(fmt.Sprintf(readQuery, table)); err != nil {
		return nil, err
	}
	if r.streamStmt, err = r.db.Reader().Prepare(fmt.Sprintf(streamQuery, table)); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *eventRepo) Read(id string) ([]AggregateEvent, error) {
	var events []AggregateEvent

	res, err := r.readStmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		event := AggregateEvent{}
		err = res.Scan(
			&event.EventType,
			&event.Event,
		)
		if err != nil {
			return events, err
		}
		events = append(events, event)
	}
	if err = res.Err(); err != nil {
		return events, err
	}
	if len(events) == 0 {
		return events, ErrNotFound
	}
	return events, nil
}

func (r *eventRepo) Stream(id string, lastSeen uint32, ch chan<- StreamResult) {
	defer close(ch)
	res, err := r.streamStmt.Query(id, lastSeen)
	if err != nil {
		ch <- StreamResult{AggregateEvent{}, err}
		return
	}
	defer res.Close()
	for res.Next() {
		event := AggregateEvent{}
		err = res.Scan(
			&event.ID,
			&event.AggregateID,
			&event.Version,
			&event.Timestamp,
			&event.EventType,
			&event.Event,
		)
		if err != nil {
			ch <- StreamResult{event, err}
			return
		}
		ch <- StreamResult{event, nil}

	}
	if err = res.Err(); err != nil {
		ch <- StreamResult{AggregateEvent{}, err}
		return
	}
}

func (r *eventRepo) Write(id string, events []AggregateEvent, version uint32) (Tx, error) {
	tx, err := r.db.Writer().Begin()
	if err != nil {
		return nil, err
	}
	insertStmt, err := tx.Prepare(r.writeStmt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, event := range events {
		version++
		result, err := insertStmt.Exec(
			event.ID,
			event.AggregateID,
			version,
			event.Timestamp,
			event.EventType,
			event.Event,
		)
		if err != nil {
			tx.Rollback()
			if pgErr, ok := err.(*pq.Error); ok && pgErr.Code.Name() == "unique_violation" {
				return nil, ErrConcurrentWrite
			}
			return nil, err
		}

		if rowsAffected, err := result.RowsAffected(); err != nil {
			tx.Rollback()
			return nil, err
		} else if rowsAffected != 1 {
			tx.Rollback()
			return nil, ErrConcurrentWrite
		}
	}
	return tx, nil
}

const readQuery = `
SELECT
event_type, event
FROM %s
WHERE aggregate_id=$1
ORDER BY version ASC`

const streamQuery = `
SELECT
id, aggregate_id, version, EXTRACT(EPOCH FROM timestamp) as timestamp, event_type, event
FROM %s
WHERE aggregate_id=$1
AND version > $2
ORDER BY version ASC`

const writeStatement = `
INSERT INTO  %s
(id, aggregate_id, version, timestamp, event_type, event)
VALUES ($1, $2, $3, TO_TIMESTAMP($4), $5, $6)`
