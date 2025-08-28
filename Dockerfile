# Stage 1: Build
FROM golang:1.24.5 AS build

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary from the web-service folder
WORKDIR /app/web-service
RUN CGO_ENABLED=0 GOOS=linux go build -o carlog

# Stage 2: Runtime
FROM debian:bookworm-slim
WORKDIR /root/

# Install certificates in case HTTPS calls are made
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
 && rm -rf /var/lib/apt/lists/*

# Copy binary from build stage
COPY --from=build /app/web-service/carlog .

CMD ["./carlog"]
