package product

import (
    "fmt"

    "github.com/marcometodo/grpc-micro-services/product-service/pkg/pb"
    "google.golang.org/grpc"
)

type ServiceClient struct {
    Client pb.ProductServiceClient
}

func InitServiceClient() pb.ProductServiceClient {
    // using WithInsecure() because no SSL running
    cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

    if err != nil {
        fmt.Println("Could not connect:", err)
    }

    return pb.NewProductServiceClient(cc)
}