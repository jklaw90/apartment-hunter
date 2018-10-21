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

	changes []events.Event
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

func (m *model) Version() uint32 {
	return m.version
}

func (m *model) GetChanges() []events.Event {
	return m.changes
}

func (m *model) apply(event events.Event, save bool) {
	switch e := event.(type) {
	case Created:
		m.address = e.Address
		m.url = e.Url
		m.title = e.Title
		m.price = e.Price
		m.bedrooms = e.Bedrooms
		m.bathrooms = e.Bathrooms
		m.sqft = e.Sqft
		m.availableDate = e.AvailableDate
		m.cats = e.Cats
		m.dogs = e.Dogs
		m.housingType = e.HousingType
		m.wdType = e.WdType
		m.parkingType = e.ParkingType
		m.images = e.Images
		m.body = e.Body
		m.lng = e.Lng
		m.lat = e.Lat
		m.closed = e.Closed

	case TitleUpdated:
		m.title = e.Title

	case PriceUpdated:
		m.price = e.Price

	case BedroomsUpdated:
		m.bedrooms = e.Bedrooms

	case BathroomsUpdated:
		m.bathrooms = e.Bathrooms

	case SqftUpdated:
		m.sqft = e.Sqft

	case AvailabilityDateUpdated:
		m.availableDate = e.AvailableDate

	case CatsUpdated:
		m.cats = e.Allowed

	case DogsUpdated:
		m.dogs = e.Allowed

	case HousingUpdated:
		m.housingType = e.HousingType

	case WDUpdated:
		m.wdType = e.WdType

	case ParkingUpdated:
		m.parkingType = e.ParkingType

	case ImagesUpdated:
		m.images = e.Images

	case BodyUpdated:
		m.body = e.Body

	case LocationUpdated:
		m.lng = e.Lng
		m.lat = e.Lat

	default:
		panic("Invalid Event")
	}

	if save {
		m.changes = append(m.changes, event)
	}
}

func Create(clID, address, url, title string, price float64,
	bedrooms uint64, bathrooms, sqft float64, availableDate string, cats,
	dogs bool, housingType, wdType, parkingType string, images []string,
	body string, lng, lat float64, closed bool) *model {
	m := &model{
		id: GetApartmentUUID(clID),
	}
	e := Created{
		Address:       address,
		Url:           url,
		Title:         title,
		Price:         price,
		Bedrooms:      bedrooms,
		Bathrooms:     bathrooms,
		Sqft:          sqft,
		AvailableDate: availableDate,
		Cats:          cats,
		Dogs:          dogs,
		HousingType:   housingType,
		WdType:        wdType,
		ParkingType:   parkingType,
		Images:        images,
		Body:          body,
		Lng:           lng,
		Lat:           lat,
		Closed:        closed,
	}
	m.apply(e, true)
	return m
}

func (m *model) TitleUpdate(title string) {
	e := TitleUpdated{
		Title: title,
	}
	m.apply(e, true)
}

func (m *model) PriceUpdate(price float64) {
	e := PriceUpdated{
		Price: price,
	}
	m.apply(e, true)
}

func (m *model) BedroomsUpdate(bedrooms uint64) {
	e := BedroomsUpdated{
		Bedrooms: bedrooms,
	}
	m.apply(e, true)
}

func (m *model) BathroomsUpdate(bathrooms float64) {
	e := BathroomsUpdated{
		Bathrooms: bathrooms,
	}
	m.apply(e, true)
}

func (m *model) SqftUpdate(sqft float64) {
	e := SqftUpdated{
		Sqft: sqft,
	}
	m.apply(e, true)
}

func (m *model) AvailabilityDateUpdate(date string) {
	e := AvailabilityDateUpdated{
		AvailableDate: date,
	}
	m.apply(e, true)
}

func (m *model) CatsUpdate(allowed bool) {
	e := CatsUpdated{
		Allowed: allowed,
	}
	m.apply(e, true)
}

func (m *model) DogsUpdate(allowed bool) {
	e := DogsUpdated{
		Allowed: allowed,
	}
	m.apply(e, true)
}

func (m *model) HousingUpdate(housingType string) {
	e := HousingUpdated{
		HousingType: housingType,
	}
	m.apply(e, true)
}

func (m *model) WDUpdate(wdType string) {
	e := WDUpdated{
		WdType: wdType,
	}
	m.apply(e, true)
}

func (m *model) ParkingUpdate(parkingType string) {
	e := ParkingUpdated{
		ParkingType: parkingType,
	}
	m.apply(e, true)
}

func (m *model) ImagesUpdate(images []string) {
	e := ImagesUpdated{
		Images: images,
	}
	m.apply(e, true)
}

func (m *model) BodyUpdate(body string) {
	e := BodyUpdated{
		Body: body,
	}
	m.apply(e, true)
}

func (m *model) LocationUpdate(lng, lat float64) {
	e := LocationUpdated{
		Lng: lng,
		Lat: lat,
	}
	m.apply(e, true)
}
