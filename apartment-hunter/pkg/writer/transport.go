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
	t.service.CreateOrUpdate(id, 0, 0, 0)
	return nil, nil
}
