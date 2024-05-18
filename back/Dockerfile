# Base image
FROM golang:1.22.1 AS base
WORKDIR /app

COPY .env hackaton-gcs.json ./

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/cosmtrek/air@latest

# Production image
FROM golang:1.22.1
WORKDIR /app

COPY --from=base /app ./
COPY --from=base /go/bin/air /usr/local/bin/air

ENTRYPOINT ["air"]
