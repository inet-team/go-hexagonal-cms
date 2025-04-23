// cmd/main.go

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Get MongoDB URI from environment or fallback to default
	mongoURI := os.Getenv("MONGODB_URL")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB connection failed:", err)
	} else {
		fmt.Println("Connected to MongoDB at", mongoURI)
	}

	app := fiber.New()

	// Serve static files from /static/images
	app.Static("/static", "./static")

	// Basic hello route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Mongo + Fiber!")
	})

	// Upload image route
	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("image")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Image file is required")
		}

		// Save the uploaded file
		savePath := fmt.Sprintf("./static/images/%s", file.Filename)
		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to save image")
		}

		imageURL := fmt.Sprintf("http://localhost:8080/static/images/%s", file.Filename)
		return c.JSON(fiber.Map{
			"message":  "Upload successful",
			"imageUrl": imageURL,
		})
	})

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
