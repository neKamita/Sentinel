# Stage 1: Build
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /notification-service -ldflags "-w -s" .

# Stage 2: Deploy
FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /notification-service /notification-service

EXPOSE 8082

ENTRYPOINT ["/notification-service"]
