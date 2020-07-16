package delivery

import (
	pb "github.com/J-Thompson12/planet-express/ship/pkg/planetexpress"
)

func GetDeliveryResponse(delivery *pb.GetDeliveryRequest) (*pb.GetDeliveryResponse, error) {
	return &pb.GetDeliveryResponse{
		Delivery: &pb.DeliveryResponse{
			Item:      delivery.Delivery.Item,
			Delivered: true,
		},
	}, nil
}

func GetDeliveryRequest(item string, quantity int32, location string) (*pb.GetDeliveryRequest, error) {
	return &pb.GetDeliveryRequest{
		Delivery: &pb.DeliveryRequest{
			Item:     item,
			Quantity: quantity,
			Location: location,
		},
	}, nil
}

func ListDeliveries(list []*pb.DeliveryResponse) (*pb.ListDeliveriesResponse, error) {
	return &pb.ListDeliveriesResponse{
		List: list,
	}, nil
}
