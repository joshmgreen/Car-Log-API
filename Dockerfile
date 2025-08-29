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

# Install Postgres client for wait-for-postgres.sh
RUN apk add --no-cache postgresql-client bash

# Copy built binary
COPY --from=build /app/carlog .

# Copy wait-for script
COPY wait-for-postgres.sh .
RUN chmod +x wait-for-postgres.sh

EXPOSE 8080

# Use wait-for-postgres.sh to start the API
CMD ["./wait-for-postgres.sh", "db", "./carlog"]
