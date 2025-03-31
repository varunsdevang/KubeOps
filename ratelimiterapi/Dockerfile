# Use a Go base image for building
FROM golang:1.24-alpine AS builder

# Set the working directory in the container
WORKDIR /app

# Copy go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . ./

# Build the Go gRPC service
RUN CGO_ENABLED=0 GOOS=linux go build -o grpc-server main.go

# Use a minimal base image for the final image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/grpc-server ./

COPY ./config/config.yml /app/config/config.yml

# Expose the gRPC port (replace with your port)
EXPOSE 3001

# Run the gRPC service
CMD ["./grpc-server"]