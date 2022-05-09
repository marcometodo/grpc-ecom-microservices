package main

import (
    "os"

    "github.com/gin-gonic/gin"
    product "github.com/marcometodo/grpc-ecom-microservices/product-service/pkg"
)

func main() {
    port := os.Getenv("HTTP_POST")

    r := gin.Default()

    product.RegisterRoutes(r)

    r.Run(port)
}