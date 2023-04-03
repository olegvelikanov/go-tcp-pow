# syntax=docker/dockerfile:1

FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/server/ cmd/server
COPY internal/app/server/ internal/app/server
COPY internal/pkg/ internal/pkg
COPY server.yaml .

RUN go build -o /server ./cmd/server

EXPOSE 3000

CMD [ "/server", "--config", "server.yaml" ]