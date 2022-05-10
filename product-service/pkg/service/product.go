package service

import (
	"context"

	pb "github.com/marcometodo/grpc-ecom-microservices/product-service/pkg/pb"
)

type Server struct {
	pb.UnimplementedProductServiceServer
}

func (s *Server) Save(ctx context.Context, in *pb.SaveRequest) (*pb.ProductResponse, error) {
	// here I can implement the save to db
	return &pb.ProductResponse{
		Id:   "1",
		Name: in.GetName(),
		Sku:  in.GetSku(),
	}, nil
}
