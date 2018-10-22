package writer

import (
	"context"

	"github.com/jklaw90/apartment-hunter/apartment-hunter/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcTransport struct {
	service Service
}

// NewGrpcTransport creates grpc transport wrapper to service
func NewGrpcTransport(s Service) *GrpcTransport {
	return &GrpcTransport{
		service: s,
	}
}

func (t *GrpcTransport) CreateOrUpdate(ctx context.Context, r *pb.CreateOrUpdateRequest) (*pb.CreateOrUpdateResponse, error) {
	if r.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	if r.GetUrl() == "" {
		return nil, status.Error(codes.InvalidArgument, "url is required")
	}

	id := GetApartmentUUID(r.GetId())
	err := t.service.CreateOrUpdate(
		id,
		r.Address,
		r.Url,
		r.Title,
		r.Price,
		r.Bedrooms,
		r.Bathrooms,
		r.Sqft,
		r.AvailableDate,
		r.Cats,
		r.Dogs,
		r.HousingType,
		r.WdType,
		r.ParkingType,
		r.Images,
		r.Body,
		r.Lng,
		r.Lat,
		false,
	)
	return &pb.CreateOrUpdateResponse{
		Id: id,
	}, err
}
