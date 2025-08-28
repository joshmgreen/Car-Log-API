# Build stage
FROM golang:1.24.5 AS build

WORKDIR /app

# Copy go.mod and go.sum first to leverage caching
COPY go.mod go.sum ./
RUN go mod download

# Copy all source files
COPY . .

# Build the binary from the cmd folder
RUN CGO_ENABLED=0 GOOS=linux go build -o carlog ./cmd/car-log-api

# Final stage
FROM alpine:latest

WORKDIR /root/
COPY --from=build /app/carlog .

EXPOSE 8080

CMD ["./carlog"]
