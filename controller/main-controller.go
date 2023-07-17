package controller

import (
	"majordome/internal/society"

	"majordome/internal/item"

	"google.golang.org/grpc"

	pb "majordome/internal/protorender/item"
)

type Controller struct {
	Society *society.Server
	Item    *item.Server
}

func (c *Controller) NewControllerSet() {
	c.Society = &society.Server{}
	c.Society.NewServer()
	c.Item = &item.Server{}
	c.Item.NewServer()
}

func (c *Controller) NewController() *grpc.Server {
	c.NewControllerSet()

	grpcServer := grpc.NewServer()

	// society.RegisterSocietyServiceServer(grpcServer, c.Society)
	pb.RegisterItemServiceServer(grpcServer, c.Item)

	return grpcServer
}
