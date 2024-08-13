# Stage 1: Build the Go application
FROM golang:1.22.5 AS builder

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy the go.mod and go.sum files for dependency management
COPY go.mod go.sum ./

# Download and verify the dependencies
RUN go mod download && go mod verify

# Copy the rest of the application source code
COPY . .

# Build the application
RUN go build -o main ./cmd/project-par

# Stage 2: Create a minimal image with the built application
FROM gcr.io/distroless/base-debian11

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy the built application from the builder stage
COPY --from=builder /usr/src/app/main .

# Expose the port on which the application runs
EXPOSE 8080

# Define the default command to run the application
CMD ["./main"]