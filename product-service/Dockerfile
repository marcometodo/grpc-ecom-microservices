FROM golang:1.18

# Port to run REST endpoint
ENV HTTP_PORT=":50051"
ENV HTTP_PORT="mongodb+srv://grpc-mongo-db:jgka7uQxhGGkthDo@cluster0.sdjlw.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
EXPOSE ${HTTP_PORT}

COPY ./product-service/ ./product-service

WORKDIR /go/product-service

RUN go build -o product-service ./cmd/main.go

ENTRYPOINT ./product-service