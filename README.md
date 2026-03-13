# HomeApp Backend

Backend API for **HomeApp**, a personal smart-home platform used to monitor and control connected devices (smart plugs, thermometers, hygrometers, etc.).

This service exposes a REST API consumed by the Flutter mobile application.

---

## Tech Stack

* **Go**
* **PostgreSQL**
* **Docker / Docker Compose**
* **JWT authentication**
* **Goose migrations**

Architecture follows a layered structure:

```
cmd/api          → application entrypoint
internal/http    → HTTP router and handlers
internal/auth    → authentication module
internal/config  → configuration loader
internal/database → database connection
migrations       → database schema migrations
```

---

## Features (current)

* User registration
* User login
* JWT authentication
* PostgreSQL persistence
* Database migrations

---

## Planned Features

* Homes management
* Rooms management
* Equipment management
* Device metrics storage
* Automation rules
* Real-time device monitoring

---

## Getting Started

### Requirements

* Go ≥ 1.22
* Docker
* Docker Compose

---

### 1. Clone the repository

```
git clone https://github.com/Revan84/HomeApp-backend.git
cd HomeApp-backend
```

---

### 2. Configure environment variables

Create a `.env` file from the example:

```
cp .env.example .env
```

Then edit the values if needed.

---

### 3. Start PostgreSQL

```
docker compose up -d
```

---

### 4. Run database migrations

```
goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/iot_flutter_app?sslmode=disable" up
```

---

### 5. Start the API

```
go run ./cmd/api
```

The API will be available at:

```
http://localhost:8080
```

---

## API Endpoints

### Health check

```
GET /health
```

---

### Authentication

Register:

```
POST /api/v1/auth/register
```

Login:

```
POST /api/v1/auth/login
```

---

## Project Status

🚧 **Work in progress**

This backend is currently under active development as part of a personal IoT home automation platform.

---

## Related Project

Flutter mobile application:

https://github.com/Revan84/HomeApp
