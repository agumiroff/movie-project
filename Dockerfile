# Start from a small, secure base image
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/server

FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary files from the builder stage
COPY --from=builder /app/app .
COPY --from=builder /app/wait-for-it.sh .
COPY --from=builder /app/internal/app/migrations ./migrations

# Expose the port that the application listens on
EXPOSE 8080

# Run the binary when the container starts
CMD ["./app"]