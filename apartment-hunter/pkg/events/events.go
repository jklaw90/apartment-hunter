package events

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
	Address       string   `json:"address"`
	Url           string   `json:"url"`
	Title         string   `json:"title"`
	Price         float64  `json:"price"`
	Bedrooms      uint64   `json:"bedrooms"`
	Bathrooms     float64  `json:"bathrooms"`
	Sqft          float64  `json:"sqft"`
	AvailableDate string   `json:"available_date"`
	Cats          bool     `json:"cats"`
	Dogs          bool     `json:"dogs"`
	HousingType   string   `json:"housing_type"`
	WdType        string   `json:"wd_type"`
	ParkingType   string   `json:"parking_type"`
	Images        []string `json:"images"`
	Body          string   `json:"body"`
	Lng           float64  `json:"lng"`
	Lat           float64  `json:"lat"`
	Closed        bool     `json:"closed"`
}

func (e Created) Type() string {
	return CreatedType
}

type TitleUpdated struct {
	Title string `json:"title"`
}

func (e TitleUpdated) Type() string {
	return TitleUpdatedType
}

type PriceUpdated struct {
	Price float64 `json:"price"`
}

func (e PriceUpdated) Type() string {
	return PriceUpdatedType
}

type BedroomsUpdated struct {
	Bedrooms uint64 `json:"bedrooms"`
}

func (e BedroomsUpdated) Type() string {
	return BedroomsUpdatedType
}

type BathroomsUpdated struct {
	Bathrooms float64 `json:"bathrooms"`
}

func (e BathroomsUpdated) Type() string {
	return BathroomsUpdatedType
}

type SqftUpdated struct {
	Sqft float64 `json:"sqft"`
}

func (e SqftUpdated) Type() string {
	return SqftUpdatedType
}

type AvailabilityDateUpdated struct {
	AvailableDate string `json:"available_date"`
}

func (e AvailabilityDateUpdated) Type() string {
	return AvailabilityDateUpdatedType
}

type CatsUpdated struct {
	Allowed bool `json:"allowed"`
}

func (e CatsUpdated) Type() string {
	return CatsUpdatedType
}

type DogsUpdated struct {
	Allowed bool `json:"allowed"`
}

func (e DogsUpdated) Type() string {
	return DogsUpdatedType
}

type HousingUpdated struct {
	HousingType string `json:"housing_type"`
}

func (e HousingUpdated) Type() string {
	return HousingUpdatedType
}

type WDUpdated struct {
	WdType string `json:"wd_type"`
}

func (e WDUpdated) Type() string {
	return WDUpdatedType
}

type ParkingUpdated struct {
	ParkingType string `json:"parking_type"`
}

func (e ParkingUpdated) Type() string {
	return ParkingUpdatedType
}

type ImagesUpdated struct {
	Images []string `json:"images"`
}

func (e ImagesUpdated) Type() string {
	return ImagesUpdatedType
}

type BodyUpdated struct {
	Body string `json:"body"`
}

func (e BodyUpdated) Type() string {
	return BodyUpdatedType
}

type LocationUpdated struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

func (e LocationUpdated) Type() string {
	return LocationUpdatedType
}
