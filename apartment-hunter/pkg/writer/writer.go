package writer

type Service interface {
	CreateOrUpdate(id string, price, lng, lat  float64) (string, error)
}

