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
}

func (e Created) Type() string {
	return CreatedType
}

type TitleUpdated struct {
}

func (e TitleUpdated) Type() string {
	return TitleUpdatedType
}

type PriceUpdated struct {
}

func (e PriceUpdated) Type() string {
	return PriceUpdatedType
}

type BedroomsUpdated struct {
}

func (e BedroomsUpdated) Type() string {
	return BedroomsUpdatedType
}

type BathroomsUpdated struct {
}

func (e BathroomsUpdated) Type() string {
	return BathroomsUpdatedType
}

type SqftUpdated struct {
}

func (e SqftUpdated) Type() string {
	return SqftUpdatedType
}

type AvailabilityDateUpdated struct {
}

func (e AvailabilityDateUpdated) Type() string {
	return AvailabilityDateUpdatedType
}

type CatsUpdated struct {
}

func (e CatsUpdated) Type() string {
	return CatsUpdatedType
}

type DogsUpdated struct {
}

func (e DogsUpdated) Type() string {
	return DogsUpdatedType
}

type HousingUpdated struct {
}

func (e HousingUpdated) Type() string {
	return HousingUpdatedType
}

type WDUpdated struct {
}

func (e WDUpdated) Type() string {
	return WDUpdatedType
}

type ParkingUpdated struct {
}

func (e ParkingUpdated) Type() string {
	return ParkingUpdatedType
}

type ImagesUpdated struct {
}

func (e ImagesUpdated) Type() string {
	return ImagesUpdatedType
}

type BodyUpdated struct {
}

func (e BodyUpdated) Type() string {
	return BodyUpdatedType
}

type LocationUpdated struct {
}

func (e LocationUpdated) Type() string {
	return LocationUpdatedType
}
