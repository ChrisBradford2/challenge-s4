## Build stage
#FROM golang:1.22.1 as builder
#WORKDIR /app
#COPY . .
#RUN go mod download
#RUN go get -u github.com/cosmtrek/air
#ENTRYPOINT ["air"]
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
#
## Final stage
#FROM alpine:latest
#RUN apk --no-cache add ca-certificates
#WORKDIR /root/
#COPY --from=builder /app/main .
#COPY .env ./
#CMD ["./main"]

# Build stage
FROM golang:1.22.1 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go install github.com/cosmtrek/air@latest

# Final stage
FROM golang:1.22.1
WORKDIR /app
COPY --from=builder /app .
COPY --from=builder /go/bin/air /usr/local/bin/air
COPY .env ./

# Set the entrypoint to run Air
ENTRYPOINT ["air"]