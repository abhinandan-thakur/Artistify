FROM golang:1.26.3 AS builder
WORKDIR /app

# Install the application dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy in the source code
COPY . .

# Build binary
RUN go build -o auth-service ./cmd/api

# Final lightweight image
FROM debian:bookworm-slim

WORKDIR /root/

COPY --from=builder /app/auth-service .

EXPOSE 8180

CMD ["./auth-service"]