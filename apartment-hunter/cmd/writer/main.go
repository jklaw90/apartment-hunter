package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jklaw90/apartment-hunter/apartment-hunter/pkg/pb"
	"github.com/jklaw90/env"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) CreateOrUpdate(ctx context.Context, in *pb.CreateOrUpdateRequest) (*pb.CreateOrUpdateResponse, error) {
	fmt.Println(in)
	return &pb.CreateOrUpdateResponse{}, nil
}

func main() {
	env := env.New(env.NewConfig(env.OptionPanicOnMissing).
		AddEnv("NAMESPACE", "67a8e8f0-2369-42a9-afd6-9b5c10f9c669", "project uuid namespace").
		AddFlagEnv("DEBUG", "debug", "false", "Enable debug logging").
		AddFlagEnv("LISTEN_ADDR", "listen", ":9000", "GRPC listen port"))

	lis, err := net.Listen("tcp", env.Get("LISTEN_ADDR"))
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
