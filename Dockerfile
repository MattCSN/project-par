# Stage 1: Build the Go application
FROM golang:1.22.5 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files for dependency management
COPY go.mod go.sum ./

# Download and verify the dependencies
RUN go mod download && go mod verify

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main cmd/project-par/main.go

# Stage 2: Create a minimal image with the built application
FROM gcr.io/distroless/base-debian11

# Set the working directory inside the container
WORKDIR /root/

# Copy the built application from the builder stage
COPY --from=builder /app/main .

# Copy the .env file if needed
COPY .env .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]