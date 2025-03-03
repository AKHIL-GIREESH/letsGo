package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	fmt.Println("Hello World")
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(200).JSON("Hello, World 👋!")
	})

	log.Fatal(app.Listen(":3000"))
}
