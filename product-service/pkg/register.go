package product

import (
    "github.com/gin-gonic/gin"
    routes "github.com/marcometodo/grpc-ecom-microservices/product-service/pkg/routes"
)

func RegisterRoutes(r *gin.Engine) *ServiceClient {
    svc := &ServiceClient{
        Client: InitServiceClient(),
    }

    routes := r.Group("/product")
    routes.POST("/", svc.Save)

    return svc
}

func (svc *ServiceClient) Save(ctx *gin.Context) {
    routes.Save(ctx, svc.Client)
}