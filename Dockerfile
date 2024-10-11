# First stage: build the application
FROM golang:1.22 AS builder

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o myapp ./cmd/service


# Expose the port (if necessary)
EXPOSE 8080

# Command to run the application
CMD ["./myapp"]
