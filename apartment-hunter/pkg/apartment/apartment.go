package apartment

import "github.com/jklaw90/apartment-hunter/apartment-hunter/pkg/hunter/data"

type ApartmentService interface {
	Get(id string) error
	ApplyEvents(events []data.AggregateEvent) error
}
