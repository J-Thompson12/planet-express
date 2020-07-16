package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/J-Thompson12/planet-express/crew"
	"github.com/J-Thompson12/planet-express/delivery"
	pb "github.com/J-Thompson12/planet-express/ship/pkg/planetexpress"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const port = 10000

type planetExpressShipServer struct {
	pb.UnimplementedPlanetExpressServer
}

func newServer() *planetExpressShipServer {
	s := &planetExpressShipServer{}
	return s
}

func (s *planetExpressShipServer) GetShip(ctx context.Context, empty *emptypb.Empty) (*pb.GetShipResponse, error) {
	crew, err := crew.GetCrew()
	if err != nil {
		return nil, err
	}

	return &pb.GetShipResponse{
		Ship: &pb.Ship{
			Name:      "planet express ship",
			Location:  "omicron persei 8",
			FuelLevel: pb.Ship_FULL,
			Crew:      crew.Crew,
		},
	}, nil
}

func (s *planetExpressShipServer) GetDelivery(ctx context.Context, deliveryRequest *pb.GetDeliveryRequest) (*pb.GetDeliveryResponse, error) {
	delivered, err := delivery.GetDeliveryResponse(deliveryRequest)
	if err != nil {
		return nil, err
	}

	return delivered, nil
}

func main() {
	log.Println("Planet Express Ship")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterPlanetExpressServer(grpcServer, newServer())

	log.Printf("Ship listening on localhost:%d\n", port)
	grpcServer.Serve(lis)
}
