# Artistify 🎵

Artistify is a backend music platform API built with Go and Gin. It supports authentication, role-based access control (RBAC), album management, Redis caching, and rate limiting.

This project is mainly being built to learn backend engineering concepts and scalable API design.

---

## Features

- JWT Authentication
- Role Based Access Control (RBAC)
- REST API using Gin
- PostgreSQL integration
- Redis caching
- Rate limiting middleware
- Album CRUD operations

---

## Tech Stack

- Go
- Gin
- PostgreSQL
- Redis
- JWT
- bcrypt

---

## API Endpoints

### Authentication

```http
POST /auth/register
POST /auth/registerWithRole
POST /auth/login