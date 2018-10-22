package writer

import (
	"github.com/satori/go.uuid"
)

type Service interface {
	CreateOrUpdate(clID, address, url, title string, price float32, bedrooms uint64,
		bathrooms, sqft float32, availableDate string, cats, dogs bool, housingType, wdType, parkingType string,
		images []string, body string, lng, lat float32, closed bool) error
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
