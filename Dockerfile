FROM golang:1.25 AS builder
WORKDIR /app

# Install the application dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy in the source code
COPY . .

# Build binary
RUN go build -o artistify ./cmd/api

# Final lightweight image
FROM debian:bookworm-slim

WORKDIR /root/

COPY --from=builder /app/artistify .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./artistify"]