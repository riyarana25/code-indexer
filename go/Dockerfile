# Use an official Golang image as a parent image
FROM golang:1.27 AS builder

# Set the working directory inside the container
WORKDIR /analyzer

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build a static Go binary (no cgo needed) so it runs on plain Alpine
RUN CGO_ENABLED=0 go build -o code-analyzer .

# Use a minimal alpine image for the final container
FROM alpine:3.22

# Set the working directory inside the container
WORKDIR /analyzer

# Copy the built binary from the builder stage
COPY --from=builder /analyzer/code-analyzer /analyzer/

# Neo4j connection settings come from .env via docker-compose

# Command to run the Go binary
CMD ["./code-analyzer"]
