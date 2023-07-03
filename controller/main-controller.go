package controller

import (
	"majordome_back/internal/society"

	"google.golang.org/grpc"

	pb "majordome_back/internal/protorender"
)

type Controller struct {
	Society *society.Server
}

func (c *Controller) NewControllerSet() {
	c.Society = &society.Server{}
	c.Society.NewServer()
}

func (c *Controller) NewController() *grpc.Server {
	c.NewControllerSet()

	grpcServer := grpc.NewServer()

	// society.RegisterSocietyServiceServer(grpcServer, c.Society)
	pb.RegisterSocietyServiceServer(grpcServer, c.Society)

	return grpcServer
}
