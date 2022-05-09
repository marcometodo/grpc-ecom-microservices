FROM golang:1.18

# Port to run REST endpoint
ENV HTTP_PORT=3000
EXPOSE ${HTTP_PORT}

COPY ./gateway-restapi/ ./gateway-restapi

WORKDIR /go/gateway-restapi

RUN go build -o gateway-restapi ./cmd/main.go

ENTRYPOINT ./gateway-restapi