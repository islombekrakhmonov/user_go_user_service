package client

import (
	"fmt"
	"user_service/config"
	"user_service/genproto/order_service"

	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	OrderService() order_service.OrderServiceClient
}

type grpcClients struct {
	orderService order_service.OrderServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	fmt.Println("NewGrpcClients")
	connOrderService, err := grpc.Dial(cfg.OrderServiceHost + cfg.OrderServicePort, grpc.WithInsecure(), 
	)
    if err != nil {
		return nil, err 
	} 

	return &grpcClients{
		orderService: order_service.NewOrderServiceClient(connOrderService),
	}, nil
}

func (g *grpcClients) OrderService() order_service.OrderServiceClient {
	return g.orderService
}


