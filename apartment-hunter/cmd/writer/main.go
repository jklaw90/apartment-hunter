package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jklaw90/apartment-hunter/apartment-hunter/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":9000"
)

type server struct{}

func (s *server) CreateOrUpdate(ctx context.Context, in *pb.CreateOrUpdateRequest) (*pb.CreateOrUpdateResponse, error) {
	fmt.Println(in)
	return &pb.CreateOrUpdateResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWriterServer(s, &server{})
	reflection.Register(s)
	log.Println("STARTING")
	defer log.Println("ENDING")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
