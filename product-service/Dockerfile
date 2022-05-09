FROM golang:1.18

# Port to run GRPC server
ENV HTTP_PORT=50051
EXPOSE ${HTTP_PORT}

WORKDIR /go/grpc-ecom-microservices

COPY ./product-service ./product-service

WORKDIR /go/grpc-ecom-microservices/product-service

RUN go build ./product-service/cmd/main.go

ENTRYPOINT ./product-service