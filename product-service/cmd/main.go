package main

import (
    "log"
    "net"
    "os"

    "github.com/marcometodo/grpc-ecom-microservices/product-service/pkg/service"
    "github.com/marcometodo/grpc-ecom-microservices/product-service/pkg/pb"
    "google.golang.org/grpc"
)

func main() {
    port := os.Getenv("HTTP_POST")

    if port == "" {
        log.Fatalf("port needs to be defined")
    }

    lis, err := net.Listen("tcp", *port)

    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    server := service.Server{}

    grpcServer := grpc.NewServer()

    pb.RegisterProductServiceServer(grpcServer, &server)
    
    log.Printf("server listening at %v", lis.Addr())

    if err := server.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}