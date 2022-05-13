package product

import (
    "context"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/marcometodo/grpc-ecom-microservices/product-service/pkg/pb"
)

func New(ctx *gin.Context, c pb.ProductServiceClient) {
    body := pb.NewRequest{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    res, err := c.New(context.Background(), &pb.NewRequest{
        Name: body.Name,
        Sku:  body.Sku,
    })

    if err != nil {
        ctx.AbortWithError(http.StatusBadGateway, err)
        return
    }

    ctx.JSON(http.StatusCreated, &res)
}