FROM golang:1.18

# Port to run GRPC server
ENV HTTP_PORT=50051
EXPOSE ${HTTP_PORT}

WORKDIR /go/grpc-micro-services

COPY ./product-service ./product-service

WORKDIR /go/grpc-micro-services/product-service

RUN go build ./product-service/cmd/main.go

ENTRYPOINT ./product-service