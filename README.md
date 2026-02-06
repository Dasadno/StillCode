# StillCode

A LeetCode-like competitive programming platform with real-time code execution in a sandboxed environment.

## Quick Start

### Prerequisites

- Docker & Docker Compose
- Go 1.23+

### 1. Start Database

```bash
docker compose up -d
```

The database is automatically initialized with schema and seeded with 35 algorithmic tasks.

### 2. Pull Code Execution Images

```bash
docker pull python:3.11-slim
docker pull gcc:12
docker pull eclipse-temurin:21-jdk-alpine
docker pull node:20-slim
docker pull golang:1.23-alpine
```

### 3. Run the Server

```bash
 go run server/cmd/main.go
```

Server starts at `http://localhost:8080`

### 4. Open the App

Navigate to `http://localhost:8080` in your browser.

## Features

- **User Authentication**: JWT-based registration and login
- **35 Algorithmic Tasks**: Arrays, Strings, DP, Trees, Graphs, and more
- **5 Languages**: Python, C++, Java, JavaScript, Go
- **Real-time Code Execution**: Docker-sandboxed runner
- **Progress Tracking**: User statistics and submissions

## Tech Stack

| Layer | Technology |
|-------|------------|
| Backend | Go 1.23 + Gin |
| Database | PostgreSQL 15 |
| Frontend | Vanilla JS + TailwindCSS |
| Editor | CodeMirror 5 |
| Runner | Docker containers |

## API Endpoints

All endpoints are **lowercase**.

### Public

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Landing page |
| GET | `/signin` | Sign in page |
| GET | `/signup` | Sign up page |
| GET | `/problems` | Problems list |
| GET | `/task/:id` | Problem solving page |
| POST | `/api/auth/signup` | Register user |
| POST | `/api/auth/signin` | Login (returns JWT) |
| GET | `/api/tasks` | List tasks |
| GET | `/api/tasks/:id` | Get task with test cases |

### Protected (require JWT)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/profile` | User profile |
| POST | `/api/run` | Run code |
| POST | `/api/submit/:id` | Submit solution |

### Query Parameters for `/api/tasks`

- `search` - Search by title
- `difficulty` - Filter: easy, medium, hard
- `community` - Filter: true/false
- `page` - Page number
- `pageSize` - Items per page

## Project Structure

```
StillCode/
├── client/
│   └── src/
│       ├── pages/          # HTML pages
│       ├── scripts/        # JavaScript modules
│       ├── services/       # API, auth, storage
│       ├── components/     # Reusable components
│       └── styles/         # CSS
├── server/
│   ├── cmd/
│   │   └── main.go         # Entry point
│   └── internal/
│       ├── auth/           # JWT
│       ├── db/             # Database
│       ├── models/         # Data models
│       ├── services/       # Business logic
│       ├── handlers/       # HTTP handlers
│       ├── middleware/     # CORS, rate limit
│       ├── router/         # Routes
│       └── runner/         # Docker sandbox
├── db/
│   ├── init.sql            # Schema
│   └── seed.sql            # 35 tasks
├── docker-compose.yml
└── README.md
```

## Code Execution

Each submission runs in an isolated Docker container:

- **Memory**: 128MB limit
- **CPU**: 0.5 cores
- **Timeout**: 5 seconds
- **Network**: Disabled
- **Filesystem**: Read-only + 64MB tmpfs

## Security

- Password hashing with bcrypt
- JWT tokens (30 min expiration)
- Rate limiting (10 req/min for code execution)
- Docker isolation with resource limits
- Input size validation

## License

MIT
