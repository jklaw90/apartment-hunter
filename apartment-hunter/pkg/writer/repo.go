package writer

import (
	"encoding/json"

	"github.com/jklaw90/apartment-hunter/apartment-hunter/pkg/hunter/data"
	"github.com/jklaw90/m2g-server/pkg/m2g/utils"
)

func (r *repository) Read(id string) (*model, error) {
	var events []Event
	aggEvents, err := r.eventRepo.Read(id)
	if err != nil {
		if err == data.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	for _, ae := range aggEvents {
		e, err := unmarshalEvent(ae.EventType, ae.Event)
		if err != nil {
			panic(err)
		}
		events = append(events, e)
	}
	return LoadModel(id, events), nil
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
			AggregateID: m.ID,
			Timestamp:   utils.Timestamp(),
			Version:     version,
			EventType:   e.Type(),
			Event:       jsonData,
		})
	}
	tx, err := r.eventRepo.Write(m.ID, aggregateEvents, m.Version())
	if err != nil {
		return err
	}
	//send events
	tx.Commit()
	return nil
}

func (r *repository) Stream(id string, lastSeen uint32, ch chan<- data.StreamResult) {
	go r.eventRepo.Stream(id, lastSeen, ch)
}
