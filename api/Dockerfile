FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the local go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the local source code to the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 3000 to the host
EXPOSE 3000

# Start the application
CMD ["./main"]