package main

import (
	"context"
	"fmt"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017/gomongodb"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	coll := client.Database("gomongodb").Collection("users")

	coll.InsertOne(context.TODO(), bson.D{
		{"name", "sherlock28"},
	})

	app := fiber.New()

	app.Use(cors.New())

	app.Static("/", "./client/dist")

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "List of users",
		})
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)
	fmt.Println("Server on port " + port)
}
