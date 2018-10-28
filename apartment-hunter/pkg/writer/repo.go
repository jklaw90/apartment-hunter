package writer

import (
	"encoding/json"

	"github.com/jklaw90/apartment-hunter/apartment-hunter/pkg/hunter/data"
	"github.com/jklaw90/apartment-hunter/apartment-hunter/pkg/hunter/utils"
	"github.com/jklaw90/apartment-hunter/apartment-hunter/pkg/writer/events"
	nats "github.com/nats-io/go-nats"
)

type repository struct {
	eventRepo  data.Repository
	eventTopic string
	nc         *nats.Conn
}

func NewRepo(eventRepo data.Repository, eventTopic string, nc *nats.Conn) *repository {
	return &repository{
		eventRepo:  eventRepo,
		eventTopic: eventTopic,
		nc:         nc,
	}
}

func (r *repository) Read(id string) (*model, error) {
	var respEvents []events.Event
	aggEvents, err := r.eventRepo.Read(id)
	if err != nil {
		if err == data.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	for _, ae := range aggEvents {
		e, err := events.UnmarshalEvent(ae.EventType, ae.Event)
		if err != nil {
			panic(err)
		}
		respEvents = append(respEvents, e)
	}

	return LoadModel(id, respEvents), nil
}

func (r *repository) Write(m *model) error {
	var aggregateEvents []data.AggregateEvent
	var version = m.Version()

	for _, e := range m.GetChanges() {
		version++
		jsonData, err := json.Marshal(e)
		if err != nil {
			return err
		}
		aggregateEvents = append(aggregateEvents, data.AggregateEvent{
			ID:          utils.NewUUID(),
			AggregateID: m.ID(),
			Timestamp:   utils.Timestamp(),
			Version:     version,
			EventType:   e.Type(),
			Event:       jsonData,
		})
	}

	tx, err := r.eventRepo.Write(m.ID(), aggregateEvents, m.Version())
	if err != nil {
		return err
	}

	events, err := json.Marshal(aggregateEvents)
	if err != nil {
		return tx.Rollback()
	}

	err = r.nc.Publish(r.eventTopic, events)
	if err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}

func (r *repository) Stream(id string, lastSeen uint32, ch chan<- data.StreamResult) {
	go r.eventRepo.Stream(id, lastSeen, ch)
}
