package writer

type model struct {
	title       string
	url         string
	housingType string
	sqft        float64
	bedrooms    uint64
	bathrooms   float64
	parkingType string
	laundryType string
	furnished   bool
	dogs        bool
	cats        bool
	images      []string
}
