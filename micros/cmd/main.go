package main

import (
	"log"

	"github.com/HsiaoCz/std-rest-api/micros"
)

func main() {
	svc := micros.NewCatFactService("https://catfact.ninja/fact")
	svc = micros.NewLoginService(svc)

	api := micros.NewApiServer(svc)
	log.Fatal(api.Start("127.0.0.1:3301"))
}
