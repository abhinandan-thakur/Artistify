# LOCAL BUILD RUN

make sure to start postegressql and redis server before...

```bash
sudo service postgresql start
```


```bash
sudo service redis-server start
```

```bash
APP_ENV=local go run cmd/api/main.go
```

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




## STATS HISTORY

###  {duration: '10s', target: 2},    { duration: '15s', target: 5} ---> 18/05/2026


  █ THRESHOLDS 

    http_req_duration
    ✓ 'p(95)<1000' p(95)=14.67ms

    http_req_failed
    ✗ 'rate<0.05' rate=42.37%


  █ TOTAL RESULTS 

    checks_total.......: 11436  457.226953/s
    checks_succeeded...: 57.61% 6589 out of 11436
    checks_failed......: 42.38% 4847 out of 11436

    ✓ status is 200
    ✗ Post status is 200
      ↳  15% — ✓ 436 / ✗ 2423
    ✗ Delete status is 200
      ↳  15% — ✓ 435 / ✗ 2424

    HTTP
    http_req_duration..............: avg=3.03ms  min=121.54µs med=668.15µs max=38.76ms p(90)=12.93ms p(95)=14.67ms
      { expected_response:true }...: avg=4.73ms  min=121.54µs med=861µs    max=38.76ms p(90)=14.41ms p(95)=15.85ms
    http_req_failed................: 42.37% 7269 out of 17154
    http_reqs......................: 17154  685.84043/s

    EXECUTION
    iteration_duration.............: avg=19.22ms min=12.11ms  med=18.39ms  max=56.79ms p(90)=24ms    p(95)=26.64ms
    iterations.....................: 2859   114.306738/s
    vus............................: 4      min=1             max=4
    vus_max........................: 5      min=5             max=5

    NETWORK
    data_received..................: 978 MB 39 MB/s
    data_sent......................: 3.8 MB 153 kB/s
