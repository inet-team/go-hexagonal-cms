# Fixing MongoDB Connection Issue in Docker Compose (Golang)

## Problem

When running a Go (Golang) application inside a Docker container, you may encounter the following error:

```
MongoDB connection failed: server selection error: context deadline exceeded
```

This happens because the Go container tries to connect to MongoDB at `localhost:27017`, which refers to itself — **not your actual host machine or another container** where MongoDB is running.

---

## ✅ Solution 1: MongoDB Running on Host Machine

If your MongoDB is running locally **outside Docker**, modify your `docker-compose.yml` to use Docker's special hostname:

```yaml
environment:
  - MONGODB_URL=mongodb://host.docker.internal:27017
```

### Notes:
- `host.docker.internal` is a special DNS name used by Docker to refer to the **host machine** (available on Windows/macOS).
- Restart your containers after making this change:

```bash
docker compose down -v
docker compose up --build
```

---

## ✅ Solution 2: MongoDB as a Docker Container

If you prefer to run MongoDB **inside Docker** alongside your Go application:

### Update `docker-compose.yml`:

```yaml
services:
  api:
    build: .
    container_name: go-hexagonal-cms-api
    environment:
      - APP_NAME=Hexagonal
      - DB_NAME_MONGODB=go-hexagonal-cms-test
      - MONGODB_URL=mongodb://mongodb:27017
    ports:
      - "8080:8080"
    volumes:
      - ./static:/app/static
    depends_on:
      - mongodb

  mongodb:
    image: mongo
    container_name: go-hexagonal-cms-mongodb
    ports:
      - "27017:27017"
```

### Update Go Code

Ensure your Go application uses the environment variable:

```go
mongoURI := os.Getenv("MONGODB_URL")
```

This allows the container to resolve `mongodb` to the correct container IP internally.

Then rebuild and start the project:

```bash
docker compose down -v
docker compose up --build
```
