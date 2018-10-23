package writer

import "fmt"

type WriterService struct {
	repo *repository
}

func New(repo *repository) *WriterService {
	return &WriterService{
		repo: repo,
	}
}

func (s *WriterService) CreateOrUpdate(uuid, address, url, title string, price float32, bedrooms uint64,
	bathrooms, sqft float32, availableDate string, cats, dogs bool, housingType, wdType, parkingType string,
	images []string, body string, lng, lat float32, closed bool) error {

	m, err := s.repo.Read(uuid)
	if err != nil {
		if err == ErrNotFound {
			m = Create(uuid, address, url, title, price,
				bedrooms, bathrooms, sqft, availableDate, cats,
				dogs, housingType, wdType, parkingType, images,
				body, lng, lat, closed)
			err = s.repo.Write(m)
			if err != nil {
				return err
			}
		}
		return err
	}

	fmt.Println(m)
	// Update if needed.
	m.TitleUpdate(title)
	m.PriceUpdate(price)
	m.BedroomsUpdate(bedrooms)
	m.BathroomsUpdate(bathrooms)
	m.SqftUpdate(sqft)
	m.AvailabilityDateUpdate(availableDate)
	m.CatsUpdate(cats)
	m.DogsUpdate(dogs)
	m.HousingUpdate(housingType)
	m.WDUpdate(wdType)
	m.ParkingUpdate(parkingType)
	m.ImagesUpdate(images)
	m.BodyUpdate(body)
	m.LocationUpdate(lng, lat)
	fmt.Println(m)

	err = s.repo.Write(m)
	if err != nil {
		return err
	}

	return nil
}
