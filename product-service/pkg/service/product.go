package service

import (
	"context"

	"github.com/marcometodo/grpc-ecom-microservices/product-service/pkg/db"
	"github.com/marcometodo/grpc-ecom-microservices/product-service/pkg/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Server struct {
	Handler db.Handler
	pb.UnimplementedProductServiceServer
}

func (server *Server) New(ctx context.Context, in *pb.NewRequest) (*pb.Product, error) {
	userCollection := db.GetCollection(server.Handler.DB)

	newProduct := pb.NewRequest{
		Name: in.GetName(),
		Sku:  in.GetSku(),
    }

	result, err := userCollection.InsertOne(ctx, newProduct)

	if err != nil {
		return &pb.Product{}, err
	}

	return &pb.Product{
		Id:   result.InsertedID.(primitive.ObjectID).Hex(),
		Name: in.GetName(),
		Sku:  in.GetSku(),
	}, nil
}
