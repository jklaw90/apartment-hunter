package writer

import "github.com/jklaw90/m2g-server/pkg/cards/events"

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

	changes []Event
	version uint32
}

func LoadModel(id string, events []events.Event) *model {
	m := &model{
		id: id,
	}
	for _, event := range events {
		m.apply(event, false)
		m.version++
	}
	return m
}

func (m *model) apply(event Event, save bool) {
	switch e := event.(type) {
	case Created:
		m.address = e.address
		m.url = e.url
		m.title = e.title
		m.price = e.price
		m.bedrooms = e.bedrooms
		m.bathrooms = e.bathrooms
		m.sqft = e.sqft
		m.availableDate = e.availableDate
		m.cats = e.cats
		m.dogs = e.dogs
		m.housingType = e.housingType
		m.wdType = e.wdType
		m.parkingType = e.parkingType
		m.images = e.images
		m.body = e.body
		m.lng = e.lng
		m.lat = e.lat
		m.closed = e.closed

	case TitleUpdated:
		m.title = e.title

	case PriceUpdated:
		m.price = e.price

	case BedroomsUpdated:
		m.bedrooms = e.bedrooms

	case BathroomsUpdated:
		m.bathrooms = e.bathrooms

	case SqftUpdated:
		m.sqft = e.sqft

	case AvailabilityDateUpdated:
		m.availableDate = e.availableDate

	case CatsUpdated:
		m.cats = e.allowed

	case DogsUpdated:
		m.dogs = e.allowed

	case HousingUpdated:
		m.housingType = e.housingType

	case WDUpdated:
		m.wdType = e.wdType

	case ParkingUpdated:
		m.parkingType = e.parkingType

	case ImagesUpdated:
		m.images = e.images

	case BodyUpdated:
		m.body = e.body

	case LocationUpdated:
		m.lng = e.lng
		m.lat = e.lat

	default:
		panic("Invalid Event")
	}

	if save {
		m.changes = append(m.changes, event)
	}
}

func Create(address string, url string, title string, price float64,
	bedrooms uint64, bathrooms float64, sqft float64, availableDate string, cats bool,
	dogs bool, housingType string, wdType string, parkingType string, images []string,
	body string, lng float64, lat float64, closed bool) *model {
	m := &model{}
	e := Created{
		address:       address,
		url:           url,
		title:         title,
		price:         price,
		bedrooms:      bedrooms,
		bathrooms:     bathrooms,
		sqft:          sqft,
		availableDate: availableDate,
		cats:          cats,
		dogs:          dogs,
		housingType:   housingType,
		wdType:        wdType,
		parkingType:   parkingType,
		images:        images,
		body:          body,
		lng:           lng,
		lat:           lat,
		closed:        closed,
	}
	m.apply(e, true)
	return m
}
