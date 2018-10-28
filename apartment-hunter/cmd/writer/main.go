package main

import (
	"log"
	"net"

	"github.com/jklaw90/apartment-hunter/apartment-hunter/pkg/hunter/data"
	"github.com/jklaw90/apartment-hunter/apartment-hunter/pkg/pb"
	"github.com/jklaw90/apartment-hunter/apartment-hunter/pkg/writer"
	"github.com/jklaw90/env"
	nats "github.com/nats-io/go-nats"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	env := env.New(env.NewConfig(env.OptionPanicOnMissing).
		AddEnv("NAMESPACE", "67a8e8f0-2369-42a9-afd6-9b5c10f9c669", "project uuid namespace").
		AddEnv("PG_WRITER", "postgresql://postgres:postgres@localhost:5432/apthunter?sslmode=disable", "postgres connection string").
		AddEnv("PG_READER", "postgresql://postgres:postgres@localhost:5432/apthunter?sslmode=disable", "postgres connection string").
		AddEnv("EVENTS_TABLE", "apartment_events", "apartments table").
		AddEnv("EVENTS_TOPIC", "apartment-events", "apartments events topic").
		AddFlagEnv("DEBUG", "debug", "false", "Enable debug logging").
		AddFlagEnv("LISTEN_ADDR", "listen", ":9000", "GRPC listen port"))

	lis, err := net.Listen("tcp", env.Get("LISTEN_ADDR"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("failed to connect to nats: %v", err)
	}

	log.Println("connecting to db")
	db := data.NewSQL(env.Get("PG_READER"), env.Get("PG_WRITER"))

	eventRepo, err := data.NewEventRepository(db, env.Get("EVENTS_TABLE"))
	if err != nil {
		panic(err)
	}

	writerRepo := writer.NewRepo(eventRepo, env.Get("EVENTS_TOPIC"), nc)

	writerService := writer.New(writerRepo)
	grpcTransort := writer.NewGrpcTransport(writerService)

	s := grpc.NewServer()
	pb.RegisterWriterServer(s, grpcTransort)
	reflection.Register(s)

	log.Println("STARTING")
	defer log.Println("ENDING")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
