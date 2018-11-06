package events

import (
	"encoding/json"
	"errors"
)

func UnmarshalEvent(eventType string, data []byte) (Event, error) {
	switch eventType {
	case CreatedType:
		e := &Created{}
		return *e, json.Unmarshal(data, e)
	case TitleUpdatedType:
		e := &TitleUpdated{}
		return *e, json.Unmarshal(data, e)
	case PriceUpdatedType:
		e := &PriceUpdated{}
		return *e, json.Unmarshal(data, e)
	case BedroomsUpdatedType:
		e := &BedroomsUpdated{}
		return *e, json.Unmarshal(data, e)
	case BathroomsUpdatedType:
		e := &BathroomsUpdated{}
		return *e, json.Unmarshal(data, e)
	case SqftUpdatedType:
		e := &SqftUpdated{}
		return *e, json.Unmarshal(data, e)
	case AvailabilityDateUpdatedType:
		e := &AvailabilityDateUpdated{}
		return *e, json.Unmarshal(data, e)
	case CatsUpdatedType:
		e := &CatsUpdated{}
		return *e, json.Unmarshal(data, e)
	case DogsUpdatedType:
		e := &DogsUpdated{}
		return *e, json.Unmarshal(data, e)
	case HousingUpdatedType:
		e := &HousingUpdated{}
		return *e, json.Unmarshal(data, e)
	case WDUpdatedType:
		e := &WDUpdated{}
		return *e, json.Unmarshal(data, e)
	case ParkingUpdatedType:
		e := &ParkingUpdated{}
		return *e, json.Unmarshal(data, e)
	case ImagesUpdatedType:
		e := &ImagesUpdated{}
		return *e, json.Unmarshal(data, e)
	case BodyUpdatedType:
		e := &BodyUpdated{}
		return *e, json.Unmarshal(data, e)
	case LocationUpdatedType:
		e := &LocationUpdated{}
		return *e, json.Unmarshal(data, e)
	}
	return nil, errors.New("event not found")
}
