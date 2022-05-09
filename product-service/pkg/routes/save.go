package product

import (
    "context"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/marcometodo/grpc-ecom-microservices/product-service/pkg/pb"
)

func Save(ctx *gin.Context, c pb.ProductServiceClient) {
    body := pb.SaveRequest{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    res, err := c.Save(context.Background(), &pb.SaveRequest{
        Name: body.Name,
        Sku:  body.Sku,
    })

    if err != nil {
        ctx.AbortWithError(http.StatusBadGateway, err)
        return
    }

    ctx.JSON(http.StatusCreated, &res)
}