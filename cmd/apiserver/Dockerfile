# Docker image of apiserver in alpine after build in multistage build
FROM golang:1.24.3 AS builder
WORKDIR /app
COPY . .
RUN ls -al /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o apiserver cmd/apiserver/main.go
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/apiserver /app/apiserver
CMD ["/app/apiserver"]
EXPOSE 8080