FROM golang:1.26.3 AS builder
WORKDIR /app

# Install the application dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy in the source code
COPY . .

# Build binary
RUN go build -o music-service ./cmd/api

# Final lightweight image
FROM debian:bookworm-slim

WORKDIR /root/

COPY --from=builder /app/music-service .

EXPOSE 8090

CMD ["./music-service"]