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

	"github.com/sherlock28/api-go-react/models"
)

const uri = "mongodb://localhost:27017/gomongodb"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Static("/", "./client/dist")

	app.Get("/users", func(c *fiber.Ctx) error {
		var users []models.User
		coll := client.Database("gomongodb").Collection("users")
		result, err := coll.Find(context.TODO(), bson.M{})

		if err != nil {
			panic(err)
		}

		for result.Next(context.TODO()) {
			var user models.User
			result.Decode(&user)
			users = append(users, user)
		}

		return c.JSON(&fiber.Map{
			"data": users,
			"msg":  "List of users",
		})
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		var user models.User

		c.BodyParser(&user)

		coll := client.Database("gomongodb").Collection("users")

		result, err := coll.InsertOne(context.TODO(), bson.D{
			{"name", user.Name},
		})

		if err != nil {
			panic(err)
		}

		return c.JSON(&fiber.Map{
			"data": result,
			"msg":  "User created successfully",
		})
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)
	fmt.Println("Server on port " + port)
}
