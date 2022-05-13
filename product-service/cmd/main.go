package main

import (
    "fmt"
    "log"
    "net"
    "os"

    "github.com/marcometodo/grpc-ecom-microservices/product-service/pkg/db"
    "github.com/marcometodo/grpc-ecom-microservices/product-service/pkg/service"
    "github.com/marcometodo/grpc-ecom-microservices/product-service/pkg/pb"
    "google.golang.org/grpc"
)

func main() {
    port := os.Getenv("PRODUCT_SERVICE_PORT")

    if port == "" {
        log.Fatalf("port needs to be defined")
    }

    handler := db.Connect()

    lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    server := service.Server{
        Handler: handler,
    }

    grpcServer := grpc.NewServer()

    pb.RegisterProductServiceServer(grpcServer, &server)
    
    log.Printf("server listening at %v", lis.Addr())

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}