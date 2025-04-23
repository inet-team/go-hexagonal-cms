# MongoDB Setup (Docker)

Quick steps to run MongoDB locally using Docker.

## 1. Pull Image

```bash
docker pull mongodb/mongodb-community-server:latest
```

## 2. Run Container

```bash
docker run --name mongodb -p 27017:27017 -d mongodb/mongodb-community-server:latest
```

## 3. Check Status

```bash
docker ps
```

## Optional: Open Shell

```bash
docker exec -it mongodb mongosh
```

> MongoDB runs on `localhost:27017` with default settings.
