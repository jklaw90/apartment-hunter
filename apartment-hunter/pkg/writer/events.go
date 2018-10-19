package writer

type Event interface {
	Type() string
}

const (
	CreatedType                 = "Created"
	TitleUpdatedType            = "TitleUpdated"
	PriceUpdatedType            = "PriceUpdated"
	BedroomsUpdatedType         = "BedroomsUpdated"
	BathroomsUpdatedType        = "BathroomsUpdated"
	SqftUpdatedType             = "SqftUpdated"
	AvailabilityDateUpdatedType = "AvailabilityDateUpdated"
	CatsUpdatedType             = "CatsUpdated"
	DogsUpdatedType             = "DogsUpdated"
	HousingUpdatedType          = "HousingUpdated"
	WDUpdatedType               = "WDUpdated"
	ParkingUpdatedType          = "ParkingUpdated"
	ImagesUpdatedType           = "ImagesUpdated"
	BodyUpdatedType             = "BodyUpdated"
	LocationUpdatedType         = "LocationUpdated"
)

type Created struct {
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

func (e Created) Type() string {
	return CreatedType
}

type TitleUpdated struct {
	title string
}

func (e TitleUpdated) Type() string {
	return TitleUpdatedType
}

type PriceUpdated struct {
	price float64
}

func (e PriceUpdated) Type() string {
	return PriceUpdatedType
}

type BedroomsUpdated struct {
	bedrooms uint64
}

func (e BedroomsUpdated) Type() string {
	return BedroomsUpdatedType
}

type BathroomsUpdated struct {
	bathrooms float64
}

func (e BathroomsUpdated) Type() string {
	return BathroomsUpdatedType
}

type SqftUpdated struct {
	sqft float64
}

func (e SqftUpdated) Type() string {
	return SqftUpdatedType
}

type AvailabilityDateUpdated struct {
	availableDate string
}

func (e AvailabilityDateUpdated) Type() string {
	return AvailabilityDateUpdatedType
}

type CatsUpdated struct {
	allowed bool
}

func (e CatsUpdated) Type() string {
	return CatsUpdatedType
}

type DogsUpdated struct {
	allowed bool
}

func (e DogsUpdated) Type() string {
	return DogsUpdatedType
}

type HousingUpdated struct {
	housingType string
}

func (e HousingUpdated) Type() string {
	return HousingUpdatedType
}

type WDUpdated struct {
	wdType string
}

func (e WDUpdated) Type() string {
	return WDUpdatedType
}

type ParkingUpdated struct {
	parkingType string
}

func (e ParkingUpdated) Type() string {
	return ParkingUpdatedType
}

type ImagesUpdated struct {
	images []string
}

func (e ImagesUpdated) Type() string {
	return ImagesUpdatedType
}

type BodyUpdated struct {
	body string
}

func (e BodyUpdated) Type() string {
	return BodyUpdatedType
}

type LocationUpdated struct {
	lng float64
	lat float64
}

func (e LocationUpdated) Type() string {
	return LocationUpdatedType
}
