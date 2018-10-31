package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/jklaw90/apartment-hunter/apartment-hunter/pkg/hunter/data"
	"github.com/jklaw90/apartment-hunter/apartment-hunter/pkg/pb"
	"github.com/jklaw90/apartment-hunter/apartment-hunter/pkg/writer/events"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/jklaw90/env"
	nats "github.com/nats-io/go-nats"
)

type server struct {
}

func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	fmt.Println(req)
	return &pb.GetResponse{}, nil
}

func main() {
	env := env.New(env.NewConfig(env.OptionPanicOnMissing).
		AddEnv("EVENTS_TOPIC", "apartment-events", "apartments events topic").
		AddEnv("EVENTS_GROUP", "apartment-group", "group for apartments events topic").
		AddFlagEnv("DEBUG", "debug", "false", "Enable debug logging").
		AddFlagEnv("LISTEN_ADDR", "listen", ":9001", "grpc listen port"))

	lis, err := net.Listen("tcp", env.Get("LISTEN_ADDR"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("failed to connect to nats: %v", err)
	}

	nc.QueueSubscribe(env.Get("EVENTS_TOPIC"), env.Get("EVENTS_GROUP"), func(m *nats.Msg) {
		var aggEvents []data.AggregateEvent
		json.Unmarshal(m.Data, &aggEvents)
		for _, aggEvt := range aggEvents {
			evts, err := events.UnmarshalEvent(aggEvt.EventType, aggEvt.Event)
			if err != nil {

			}
			fmt.Printf("Received a message: %v\n", evts)
		}
	})

	gserver := &server{}
	s := grpc.NewServer()
	pb.RegisterApartmentServer(s, gserver)
	reflection.Register(s)

	log.Println("STARTING")
	defer log.Println("ENDING")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
