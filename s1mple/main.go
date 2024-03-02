package main

import (
	"flag"

	"github.com/HsiaoCz/std-rest-api/s1mple/api"
	"github.com/gofiber/fiber/v2"
)

// simple
// some detail of the goddamn coding
func main() {
	listenAddr := flag.String("listenAddr", "127.0.0.1:8001", "set the server listen address")
	flag.Parse()
	app := fiber.New()
	v1 := app.Group("/api/v1")
	{
		v1.Post("/user", api.HandleCreateUser)
	}
	app.Listen(*listenAddr)
}
