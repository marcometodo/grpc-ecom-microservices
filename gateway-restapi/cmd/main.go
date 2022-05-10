package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    product "github.com/marcometodo/grpc-ecom-microservices/product-service/pkg"
)

func main() {
    port := os.Getenv("HTTP_PORT")

    if port == "" {
        log.Fatalf("port needs to be defined")
    }

    r := gin.Default()

    product.RegisterRoutes(r)

    r.Run(port)
}