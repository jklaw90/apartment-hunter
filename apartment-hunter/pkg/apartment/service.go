package apartment

import "github.com/jklaw90/apartment-hunter/apartment-hunter/pkg/hunter/data"

type Service struct {
}

func (s *Service) Get(id string) error {
	return nil
}

func (s *Service) ApplyEvents(events []data.AggregateEvent) error {
	return nil
}
