# Stage 1: Build
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
# -ldflags "-w -s" reduces the size of the binary
# CGO_ENABLED=0 is important for a truly static binary
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /ingestor-gateway -ldflags "-w -s" .

# Stage 2: Deploy
FROM alpine:latest

# Add ca-certificates to make SSL calls
RUN apk --no-cache add ca-certificates

# Copy the built binary from the builder stage
COPY --from=builder /ingestor-gateway /ingestor-gateway

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/ingestor-gateway"]
