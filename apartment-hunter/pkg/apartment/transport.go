package apartment

import (
	"context"

	"github.com/jklaw90/apartment-hunter/apartment-hunter/pkg/pb"
)

type GrpcTransport struct {
}

func (s *GrpcTransport) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{}, nil
}
