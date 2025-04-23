# Golang Hexagonal Architecture


## Initialize a Go Module

```bash
go mod init go-hexagonal-cms

# Install fiber dependencies
go get github.com/gofiber/fiber/v2

# Install mongodb dependencies
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
```


## Project Structure

```
_docs
    - setup-mongodb-docker.md
cmd/
    - main.go
go.mod
go.sum
```


## Run App With Go

```bash
go run cmd/main.go
```
