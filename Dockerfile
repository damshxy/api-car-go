# Use the official Golang image with the correct version as a base image
FROM golang:1.23-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code to the Working Directory
COPY . .

# Build the Go app
RUN go build -o main cmd/main.go

# Expose port 5050 to the outside world
EXPOSE 5050

# Command to run the executable
CMD ["./main"]