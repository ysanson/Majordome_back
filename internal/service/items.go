package service

import (
	"context"

	pb "majordome/api/v1/tracker"
)

type ItemsService struct {
	pb.UnimplementedItemsServer
}

func NewItemsService() *ItemsService {
	return &ItemsService{}
}

func (s *ItemsService) GetItems(ctx context.Context, req *pb.GetItemsRequest) (*pb.GetItemsResponse, error) {
	return &pb.GetItemsResponse{}, nil
}
func (s *ItemsService) GetItemsStream(req *pb.GetItemsRequest, conn pb.Items_GetItemsStreamServer) error {
	for {
		err := conn.Send(&pb.Item{})
		if err != nil {
			return err
		}
	}
}
func (s *ItemsService) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.Item, error) {
	return &pb.Item{}, nil
}
func (s *ItemsService) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.Item, error) {
	return &pb.Item{}, nil
}
func (s *ItemsService) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.Item, error) {
	return &pb.Item{}, nil
}
func (s *ItemsService) DeleteItem(ctx context.Context, req *pb.GetItemRequest) (*pb.Item, error) {
	return &pb.Item{}, nil
}
func (s *ItemsService) NewRealisation(ctx context.Context, req *pb.GetItemRequest) (*pb.Realisation, error) {
	return &pb.Realisation{}, nil
}
