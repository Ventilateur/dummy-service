#### Build stage
FROM golang:1.14-alpine as builder

WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download
COPY . /app/

RUN CGO_ENABLED=0 GOOS=linux go build -o sample-svc

#### Runtime
FROM alpine:3.12
EXPOSE 8080

WORKDIR /app
COPY --from=builder /app/sample-svc /app/

ENTRYPOINT ["./sample-svc"]
