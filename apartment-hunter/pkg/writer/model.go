package writer

type model struct {
	id            string
	address       string
	url           string
	title         string
	price         float64
	bedrooms      uint64
	bathrooms     float64
	sqft          float64
	availableDate string
	cats          bool
	dogs          bool
	housingType   string
	wdType        string
	parkingType   string
	images        []string
	body          string
	lng           float64
	lat           float64
	closed        bool
}
