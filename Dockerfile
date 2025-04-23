# Stage 0 - Building server application
FROM golang:1.22.4 as builder
WORKDIR /app
ADD . .

# Disable CGO and compile server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w' -o server ./cmd/main.go
RUN chmod +x server

# Stage 1 - Server start
FROM --platform=amd64 alpine:3.9
WORKDIR /app

# Install timezone data
RUN apk add --no-cache tzdata

# Set the timezone
ENV TZ=Asia/Bangkok

# copy binary
COPY --from=builder /app/server /app/server

# Copy the static files to the container
COPY static/files /app/static/files
COPY static/images /app/static/images

# Expose the same port (8080) for container and host
EXPOSE 8080
CMD [ "/app/server", "start" ]
