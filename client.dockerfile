# syntax=docker/dockerfile:1

FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/client/ cmd/client
COPY internal/app/client/ internal/app/client
COPY internal/pkg/ internal/pkg

RUN go build -o /client ./cmd/client

EXPOSE 3000

CMD ["/client"]