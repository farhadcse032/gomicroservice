# Use official Golang image
FROM golang:1.22.6-alpine

# Set work directory
WORKDIR /app

# Copy source code
COPY . .

# Build the Go service
RUN go mod tidy && go build -o service3

# Run the service
CMD ["./service3"]
