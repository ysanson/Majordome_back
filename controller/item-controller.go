package controller

import (
	"majordome/internal/item"

	"google.golang.org/grpc"

	pb "majordome/internal/protorender/item"
)

type ItemController struct {
	Item *item.Server
}

func (c *ItemController) NewControllerSet() {
	c.Item = &item.Server{}
	c.Item.NewServer()
}

func (c *ItemController) NewController() *grpc.Server {
	c.NewControllerSet()

	grpcServer := grpc.NewServer()

	// society.RegisterSocietyServiceServer(grpcServer, c.Society)
	pb.RegisterSocietyServiceServer(grpcServer, c.Society)

	
	return grpcServer
}
