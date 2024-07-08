# Use golang as the base image
FROM golang:1.16

# Set the current working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o server ./server/main.go
RUN go build -o client ./client/main.go

# Expose port 50051 to the outside world
EXPOSE 50051

# Command to run the gRPC server
CMD ["./server"]
