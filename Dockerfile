# Use the official Go image as the base image
FROM golang:1.22.5

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

# Expose the port on which the application runs
EXPOSE 8080

# Define the default command to run the application
CMD ["./main"]