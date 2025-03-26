# Use the official Golang image for building
FROM golang:1.23 AS builder
# Set working directory
WORKDIR /app
# Copy Go modules and dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy source code
COPY . .
# Build the application
RUN go build -o gosearch .

# Use a minimal base image for final deployment
FROM alpine:latest
# Set working directory in the container
WORKDIR /root/
# Copy the built binary from the builder stage
COPY --from=builder /app/gosearch .
# Expose the application port
EXPOSE 3000
# Run the application
CMD ["./gosearch"]