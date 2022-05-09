package main

import (
    "os"

    "github.com/gin-gonic/gin"
    product "github.com/marcometodo/grpc-ecom-microservices/product-service/pkg"
)

func main() {
    port := os.Getenv("HTTP_POST")

    if port == "" {
        port = ":3000"
    }

    r := gin.Default()

    product.RegisterRoutes(r)

    r.Run(port)
}