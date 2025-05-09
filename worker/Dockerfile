# Worker Dockerfile
FROM golang:1.20-alpine

WORKDIR /app

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum* ./
RUN go mod download

# Then copy the rest of the source code
COPY . .

# Build the application
RUN go build -o worker main.go

# Run with worker mode
CMD ["./worker", "-mode=worker"]