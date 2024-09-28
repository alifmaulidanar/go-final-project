# Use the official Golang image for building the application
FROM golang:1.22.5 AS build

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed.
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Start a new stage from scratch
FROM golang:1.22.5

# Set the working directory
WORKDIR /app

# Copy the pre-built binary from the build stage
COPY --from=build /app/main .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
