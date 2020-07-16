package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/J-Thompson12/planet-express/delivery"
	"github.com/golang/protobuf/ptypes/empty"

	pb "github.com/J-Thompson12/planet-express/ship/pkg/planetexpress"

	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)

func getShip(client pb.PlanetExpressClient) (pb.Ship, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetShip(ctx, &empty.Empty{})

	if err != nil {
		log.Fatalf("%v.GetShip(_) = _, %v: ", client, err)
		return pb.Ship{}, err
	}

	return *resp.Ship, nil
}

func getDelivery(client pb.PlanetExpressClient) (pb.DeliveryResponse, error) {
	delivery, err := delivery.GetDeliveryRequest("Holophonor", 2, "Mars - Wongs")
	if err != nil {
		return pb.DeliveryResponse{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetDelivery(ctx, delivery)

	if err != nil {
		log.Fatalf("%v.GetShip(_) = _, %v: ", client, err)
		return pb.DeliveryResponse{}, err
	}

	return *resp.Delivery, nil
}

func main() {
	log.Println("Planet Express Headquarters")

	flag.Parse()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewPlanetExpressClient(conn)
	log.Printf("Connected to planet express ship with addr: %s\n\n", *serverAddr)

	list := []*pb.DeliveryResponse{}

	ship, _ := getShip(client)
	deliveryResponse, _ := getDelivery(client)
	list = append(list, &deliveryResponse)
	listDeliveries, err := delivery.ListDeliveries(list)

	ship.List = listDeliveries.List

	jsonString, err := json.Marshal(ship)
	ioutil.WriteFile("../planet-express.json", jsonString, os.ModePerm)
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Println(ship)
}
