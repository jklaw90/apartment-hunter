package writer

import (
	"context"

	"github.com/jklaw90/apartment-hunter/apartment-hunter/pkg/pb"
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
	return nil, nil
}
