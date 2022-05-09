package product

import (
	"context"
	"log"
	"net"
	"flag"
	"fmt"

	"google.golang.org/grpc"
	"github.com/marcometodo/grpc-ecom-microservices/product-service/pkg/pb"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedProductServiceServer
}

func (s *server) Save(ctx context.Context, in *pb.SaveRequest) (*pb.ProductResponse, error) {
	// here I can implement the save to db
	return &pb.ProductResponse{
		Id:   "1",
		Name: in.GetName(),
		Sku:  in.GetSku(),
	}, nil
}

func StartgRPCServer() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

    if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
