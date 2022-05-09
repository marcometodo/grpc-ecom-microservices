FROM golang:1.18

# Port to run REST endpoint
ENV HTTP_PORT=":3000"
EXPOSE ${HTTP_PORT}

COPY ./product-service/ ./product-service

WORKDIR /go/product-service

RUN go build -o product-service ./cmd/main.go

ENTRYPOINT ./product-service