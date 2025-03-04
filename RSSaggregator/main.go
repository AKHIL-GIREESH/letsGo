package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{}
var nextID = 1

func main() {
	fmt.Println("Hello World")
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(200).JSON(users)
	})

	app.Post("/users", func(c fiber.Ctx) error {
		user := new(User)
		if err := c.Bind().JSON(user); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		user.ID = nextID
		nextID++
		users = append(users, *user)
		return c.Status(200).JSON(user)
	})

	app.Patch("/users/:id", func(c fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		var user *User = nil

		for i := range users {
			if users[i].ID == id {
				user = &users[i]
				break
			}
		}

		if user == nil {
			return c.Status(404).JSON(fiber.Map{"error": "User not found"})
		}

		var updates User
		if err := c.Bind().JSON(&updates); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if updates.Name != "" {
			user.Name = updates.Name
		}

		return c.Status(200).JSON(user)
	})

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
