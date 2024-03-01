package main

import (
	"flag"

	"github.com/HsiaoCz/std-rest-api/hotel/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// parse the command line address
	// some command like this:
	// make run --listenAddr : 7000
	listenAddr := flag.String("listenAddr", "127.0.0.1:9001", "set up the server listen address")
	flag.Parse()
	app := fiber.New()
	v1 := app.Group("/api/v1")
	v1.Get("/user", handlers.HandleGetUser)
	v1.Get("/user/:id",)
	app.Listen(*listenAddr)
}
