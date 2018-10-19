package writer

import (
	"github.com/satori/go.uuid"
)

type Service interface {
	CreateOrUpdate(id string, price, lng, lat float64) (string, error)
}

const (
	CRAIGSLIST_NS = "c0e173e3-b11b-48d1-aa82-554d8505acd5"
)

// GetApartmentUUID Used to convert CL id to our internal id
func GetApartmentUUID(clID string) string {
	clNs, err := uuid.FromString(CRAIGSLIST_NS)
	if err != nil {
		panic(err)
	}
	return uuid.NewV5(clNs, clID).String()
}
