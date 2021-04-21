#### Build stage
FROM golang:1.16-alpine as builder

WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download
COPY . /app/

RUN CGO_ENABLED=0 GOOS=linux go build -o dummy-service

#### Runtime
FROM alpine:3.13
EXPOSE 8080

WORKDIR /app
COPY --from=builder /app/dummy-service /app/

ENTRYPOINT ["./dummy-service"]
