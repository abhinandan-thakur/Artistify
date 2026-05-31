# Artistify 🎵

A backend music platform API built with **Go**, designed around a **microservices architecture**. Artistify supports JWT authentication, role-based access control (RBAC), album/music management, Redis caching, rate limiting, and full observability via Prometheus and Grafana.

> 🚧 This project is actively being built to learn and explore backend engineering concepts, scalable API design, and microservices patterns.

---

## Table of Contents

- [Architecture](#architecture)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Run with Docker (Recommended)](#run-with-docker-recommended)
  - [Run Locally](#run-locally)
- [Environment Variables](#environment-variables)
- [API Endpoints](#api-endpoints)
- [Monitoring](#monitoring)
- [Load Testing](#load-testing)

---

## Architecture

Artistify follows a **microservices** pattern with the following services:

| Service                | Description                                               | Port(s)          |
|------------------------|-----------------------------------------------------------|------------------|
| `auth-service`         | Handles user registration, login, and JWT issuance        | HTTP: 8180, gRPC: 50051 |
| `music-service`        | Manages albums, tracks, and music catalog operations      | HTTP: 8080       |
| `notification-service` | Handles email/notification delivery                       | —                |
| `prometheus`           | Metrics collection from all services                      | 9090             |
| `grafana`              | Metrics visualization dashboards                          | 3000             |

Services communicate internally via **gRPC** (defined in `/proto`). PostgreSQL is the primary datastore, with Redis used for caching.

---

## Features

- **JWT Authentication** — Secure token-based auth with configurable secrets
- **Role-Based Access Control (RBAC)** — Fine-grained permissions per user role
- **Microservices Architecture** — Independently deployable services with gRPC inter-service communication
- **REST API with Gin** — Fast and lightweight HTTP routing
- **PostgreSQL Integration** — Persistent data storage via `pgx/v5`
- **Redis Caching** — Response caching to reduce DB load
- **Rate Limiting Middleware** — Protect endpoints from abuse
- **Album / Music CRUD** — Full create, read, update, and delete for the music catalog
- **Prometheus Metrics** — Instrumented endpoints with `prometheus/client_golang`
- **Grafana Dashboards** — Visualise service health and request metrics
- **k6 Load Testing** — Performance test scripts in `/tests/k6`
- **GitHub Actions CI/CD** — Automated build and deployment pipeline
- **Docker Compose** — One-command local environment setup

---

## Tech Stack

| Layer          | Technology                              |
|----------------|-----------------------------------------|
| Language       | Go 1.25                                 |
| HTTP Framework | Gin                                     |
| Database       | PostgreSQL 17 (via `pgx/v5`)            |
| Cache          | Redis 8                                 |
| Auth           | JWT (`golang-jwt/jwt v5`) + bcrypt      |
| RPC            | gRPC + Protocol Buffers                 |
| Observability  | Prometheus + Grafana                    |
| Load Testing   | k6                                      |
| Containerization | Docker + Docker Compose               |
| CI/CD          | GitHub Actions                          |

---

## Project Structure

```
Artistify/
├── auth-service/          # Authentication microservice (HTTP + gRPC)
├── music-service/         # Music catalog microservice
├── notification-service/  # Notification microservice
├── proto/                 # Protobuf definitions for gRPC
├── monitoring/            # Prometheus configuration
├── templates/             # Email/HTML templates
├── tests/k6/              # k6 load test scripts
├── .github/workflows/     # CI/CD pipelines
├── docker-compose.yml     # Full local stack definition
├── .env.docker.example    # Env vars template for Docker setup
├── .env.local.example     # Env vars template for local setup
└── go.mod                 # Go module definition
```

---

## Getting Started

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/) & [Docker Compose](https://docs.docker.com/compose/) (for Docker setup)
- [Go 1.25+](https://go.dev/dl/) (for local setup)
- PostgreSQL & Redis (for local setup only)

---

### Run with Docker (Recommended)

This spins up all services — Postgres, Redis, auth-service, music-service, Prometheus, and Grafana — with a single command.

**1. Clone the repository**

```bash
git clone https://github.com/abhinandan-thakur/Artistify.git
cd Artistify
```

**2. Set up environment files**

Each service needs its own `.env.compose` file. Use the provided example as a starting point:

```bash
cp .env.docker.example auth-service/.env.compose
cp .env.docker.example music-service/.env.compose
```

Edit the values as needed (see [Environment Variables](#environment-variables)).

**3. Start all services**

```bash
docker compose up --build
```

| Service      | URL                     |
|--------------|-------------------------|
| Music API    | http://localhost:8080   |
| Auth API     | http://localhost:8180   |
| Prometheus   | http://localhost:9090   |
| Grafana      | http://localhost:3000   |

---

### Run Locally

**1. Start required infrastructure**

```bash
sudo service postgresql start
sudo service redis-server start
```

**2. Set up environment**

```bash
cp .env.local.example .env
# Edit .env with your local DB credentials
```

**3. Run the desired service**

```bash
# Auth service
cd auth-service
APP_ENV=local go run cmd/api/main.go

# Music service
cd music-service
APP_ENV=local go run cmd/api/main.go
```

---

## Environment Variables

| Variable              | Description                          | Example                  |
|-----------------------|--------------------------------------|--------------------------|
| `DB_HOST`             | PostgreSQL host                      | `localhost` / `postgres` |
| `DB_PORT`             | PostgreSQL port                      | `5432`                   |
| `DB_USER`             | PostgreSQL username                  | `postgres`               |
| `DB_NAME`             | PostgreSQL database name             | `musicdb`                |
| `DB_PASSWORD`         | PostgreSQL password                  | `your-password`          |
| `URL_PORT`            | HTTP server port                     | `8080`                   |
| `REDIS_HOST`          | Redis host                           | `localhost` / `redis`    |
| `REDIS_PORT`          | Redis port                           | `6379`                   |
| `JWT_SECRET`          | Secret key for signing JWTs          | `super-secret-key`       |
| `GRPC_HOST`           | gRPC server host (auth-service)      | `auth-service`           |
| `GRPC_PORT`           | gRPC server port                     | `50051`                  |

See `.env.docker.example` and `.env.local.example` for a full reference.

---

## API Endpoints

### Auth Service (`localhost:8180`)

| Method | Endpoint                  | Description                          | Auth Required |
|--------|---------------------------|--------------------------------------|---------------|
| POST   | `/auth/register`          | Register a new user                  | No            |
| POST   | `/auth/registerWithRole`  | Register a user with a specific role | No            |
| POST   | `/auth/login`             | Login and receive a JWT              | No            |

### Music Service (`localhost:8080`)

Albums and track management endpoints are served here. JWT token from the auth service is required for protected routes.

---

## Monitoring

Prometheus scrapes metrics from both `auth-service` and `music-service`. Grafana visualises these on port `3000`.

To access Grafana after running Docker Compose:
1. Open http://localhost:3000
2. Default credentials: `admin` / `admin`
3. Add Prometheus as a data source (`http://prometheus:9090`)

---

## Load Testing

k6 scripts are located in `/tests/k6`. To run a load test:

```bash
k6 run tests/k6/<script-name>.js
```

**Example results** (2-stage ramp: 10s @ 2 VUs → 15s @ 5 VUs):

```
http_req_duration   avg=3.03ms  p(95)=14.67ms
http_reqs           17154       685/s
data_received       978 MB      39 MB/s
```

---

## Generating Proto Files

If you modify `.proto` definitions in `/proto`, regenerate the Go stubs with:

```bash
protoc \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  proto/<your-file>.proto
```

---

*Built with ❤️ to learn Go, microservices, and scalable API design.*
