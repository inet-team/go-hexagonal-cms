# Golang Hexagonal Architecture


## Initialize a Go Module

```bash
go mod init go-hexagonal-cms

# Install fiber
go get github.com/gofiber/fiber/v2

# Install MongoDB driver
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
```

## Project Structure

```
_docs/
    - docker/
cmd/
    - main.go
static/
    - files/
        - example-image.webp
    - images/
        - example-image.webp
Dockerfile
docker-compose.yml
go.mod
go.sum
README.md
```

## Run with Docker Compose

```bash
docker compose down -v

docker compose up --build
```

## Run the App with Directly (Not recommended)

```bash
go run cmd/main.go
```
