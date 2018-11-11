package apartment

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jklaw90/apartment-hunter/apartment-hunter/pkg/apartment/events"
	"github.com/jklaw90/apartment-hunter/apartment-hunter/pkg/hunter/data"
	nats "github.com/nats-io/go-nats"
)

type Subscriber struct {
	conn *nats.Conn
}

func NewSubscriber(url string) *Subscriber {
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatalf("failed to connect to nats: %v", err)
	}

	return &Subscriber{
		conn: nc,
	}
}
func (s *Subscriber) SubGroup(topic, group string, callback func([]data.AggregateEvent) error) {
	s.conn.QueueSubscribe(topic, group, func(m *nats.Msg) {
		var aggEvents []data.AggregateEvent
		json.Unmarshal(m.Data, &aggEvents)
		for _, aggEvt := range aggEvents {
			evts, err := events.UnmarshalEvent(aggEvt.EventType, aggEvt.Event)
			if err != nil {

			}
			fmt.Printf("Received a message: %v\n", evts)
		}
	})
}

func (s *Subscriber) Close() {

}

func (s *Subscriber) HandleAggregateEvents(events []data.AggregateEvent) error {
	return nil
}
