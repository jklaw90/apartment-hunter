package writer

import (
	"github.com/jklaw90/apartment-hunter/apartment-hunter/pkg/writer/events"
)

type model struct {
	id            string
	address       string
	url           string
	title         string
	price         float32
	bedrooms      uint64
	bathrooms     float32
	sqft          float32
	availableDate string
	cats          bool
	dogs          bool
	housingType   string
	wdType        string
	parkingType   string
	images        []string
	body          string
	lng           float32
	lat           float32
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

func (m *model) ID() string {
	return m.id
}

func (m *model) Version() uint32 {
	return m.version
}

func (m *model) GetChanges() []events.Event {
	return m.changes
}

func (m *model) apply(event events.Event, save bool) {
	switch e := event.(type) {
	case events.Created:
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

	case events.TitleUpdated:
		m.title = e.Title

	case events.PriceUpdated:
		m.price = e.Price

	case events.BedroomsUpdated:
		m.bedrooms = e.Bedrooms

	case events.BathroomsUpdated:
		m.bathrooms = e.Bathrooms

	case events.SqftUpdated:
		m.sqft = e.Sqft

	case events.AvailabilityDateUpdated:
		m.availableDate = e.AvailableDate

	case events.CatsUpdated:
		m.cats = e.Allowed

	case events.DogsUpdated:
		m.dogs = e.Allowed

	case events.HousingUpdated:
		m.housingType = e.HousingType

	case events.WDUpdated:
		m.wdType = e.WdType

	case events.ParkingUpdated:
		m.parkingType = e.ParkingType

	case events.ImagesUpdated:
		m.images = e.Images

	case events.BodyUpdated:
		m.body = e.Body

	case events.LocationUpdated:
		m.lng = e.Lng
		m.lat = e.Lat

	default:
		panic("Invalid Event")
	}

	if save {
		m.changes = append(m.changes, event)
	}
}

func Create(id, address, url, title string, price float32,
	bedrooms uint64, bathrooms, sqft float32, availableDate string, cats,
	dogs bool, housingType, wdType, parkingType string, images []string,
	body string, lng, lat float32, closed bool) *model {
	m := &model{
		id: id,
	}
	e := events.Created{
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
	if m.title == title {
		return
	}
	e := events.TitleUpdated{
		Title: title,
	}
	m.apply(e, true)
}

func (m *model) PriceUpdate(price float32) {
	if m.price == price {
		return
	}
	e := events.PriceUpdated{
		Price: price,
	}
	m.apply(e, true)
}

func (m *model) BedroomsUpdate(bedrooms uint64) {
	if m.bedrooms == bedrooms {
		return
	}
	e := events.BedroomsUpdated{
		Bedrooms: bedrooms,
	}
	m.apply(e, true)
}

func (m *model) BathroomsUpdate(bathrooms float32) {
	if m.bathrooms == bathrooms {
		return
	}
	e := events.BathroomsUpdated{
		Bathrooms: bathrooms,
	}
	m.apply(e, true)
}

func (m *model) SqftUpdate(sqft float32) {
	if m.sqft == sqft {
		return
	}
	e := events.SqftUpdated{
		Sqft: sqft,
	}
	m.apply(e, true)
}

func (m *model) AvailabilityDateUpdate(date string) {
	if m.availableDate == date {
		return
	}
	e := events.AvailabilityDateUpdated{
		AvailableDate: date,
	}
	m.apply(e, true)
}

func (m *model) CatsUpdate(allowed bool) {
	if m.cats == allowed {
		return
	}
	e := events.CatsUpdated{
		Allowed: allowed,
	}
	m.apply(e, true)
}

func (m *model) DogsUpdate(allowed bool) {
	if m.dogs == allowed {
		return
	}
	e := events.DogsUpdated{
		Allowed: allowed,
	}
	m.apply(e, true)
}

func (m *model) HousingUpdate(housingType string) {
	if m.housingType == housingType {
		return
	}
	e := events.HousingUpdated{
		HousingType: housingType,
	}
	m.apply(e, true)
}

func (m *model) WDUpdate(wdType string) {
	if m.wdType == wdType {
		return
	}
	e := events.WDUpdated{
		WdType: wdType,
	}
	m.apply(e, true)
}

func (m *model) ParkingUpdate(parkingType string) {
	if m.parkingType == parkingType {
		return
	}
	e := events.ParkingUpdated{
		ParkingType: parkingType,
	}
	m.apply(e, true)
}

func (m *model) ImagesUpdate(images []string) {
	if len(m.images) == len(images) {
		if len(images) == 0 {
			return
		}
		var dupe = false
		for i, img := range m.images {
			if img != images[i] {
				dupe = false
				break
			} else {
				dupe = true
			}
		}
		if dupe {
			return
		}
	}

	e := events.ImagesUpdated{
		Images: images,
	}
	m.apply(e, true)
}

func (m *model) BodyUpdate(body string) {
	if m.body == body {
		return
	}
	e := events.BodyUpdated{
		Body: body,
	}
	m.apply(e, true)
}

func (m *model) LocationUpdate(lng, lat float32) {
	if m.lng == lng && m.lat == lat {
		return
	}
	e := events.LocationUpdated{
		Lng: lng,
		Lat: lat,
	}
	m.apply(e, true)
}
