package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/foo", handleFoo)
	app.Listen("127.0.0.1:9001")
}

func handleFoo(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(map[string]string{"msg": "working just"})
}
