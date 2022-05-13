package product

import (
    "fmt"
    "os"

    "github.com/marcometodo/grpc-ecom-microservices/product-service/pkg/pb"
    "google.golang.org/grpc"
)

type ServiceClient struct {
    Client pb.ProductServiceClient
}

func InitServiceClient() pb.ProductServiceClient {
    port := os.Getenv("PRODUCT_SERVICE_PORT")

    if port == "" {
        log.Fatalf("pgrc port needs to be defined")
    }

    ip := os.Getenv("PRODUCT_SERVICE_IP")

    if ip == "" {
        log.Fatalf("grpc IP needs to be defined")
    }

    // using WithInsecure() because no SSL running
    cc, err := grpc.Dial(fmt.Sprintf("%s:%s", ip, port), grpc.WithInsecure())

    if err != nil {
        fmt.Println("Could not connect:", err)
    }

    return pb.NewProductServiceClient(cc)
}