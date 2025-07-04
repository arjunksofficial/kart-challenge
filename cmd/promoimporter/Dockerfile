# Docker image of promoimporter in alpine after build in multistage build
FROM golang:1.24.3 AS builder
WORKDIR /app
COPY . .
RUN ls -al /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o promoimporter cmd/promoimporter/main.go
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/promoimporter /app/promoimporter
CMD ["/app/promoimporter"]
EXPOSE 8080