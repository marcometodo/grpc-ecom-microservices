package service

import (
	"context"

	"github.com/marcometodo/grpc-ecom-microservices/product-service/pkg/db"
	"github.com/marcometodo/grpc-ecom-microservices/product-service/pkg/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Server struct {
	Handler db.Handler
}

func (server *Server) Save(ctx context.Context, in *pb.SaveRequest) (*pb.ProductResponse, error) {
	userCollection := db.GetCollection()

	newProduct := pb.SaveRequest{
        Id:   primitive.NewObjectID(),
		Name: in.GetName(),
		Sku:  in.GetSku(),
    }

	result, err := userCollection.InsertOne(ctx, newProduct)

	if err != nil {
		return &pb.ProductResponse{}, err
	}

	return &pb.ProductResponse{
		Id:   "1",
		Name: in.GetName(),
		Sku:  in.GetSku(),
	}, nil
}
