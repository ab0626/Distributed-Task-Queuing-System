# API Dockerfile
FROM golang:1.24-alpine

WORKDIR /app

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum* ./
RUN go mod download

# Then copy the rest of the source code
COPY . .

# Build the application
RUN go build -o api-server main.go

# Run the API server with appropriate flags
CMD ["./api-server", "-mode=api"]