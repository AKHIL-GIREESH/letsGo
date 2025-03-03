package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
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

	log.Fatal(app.Listen(":3000"))
}
