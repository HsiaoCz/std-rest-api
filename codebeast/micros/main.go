package main

import (
	"log"

	"github.com/HsiaoCz/std-rest-api/codebeast/micros/api"
	"github.com/HsiaoCz/std-rest-api/codebeast/micros/service"
)

func main() {
	svc := service.NewCatFactService("https://catfact.ninja/fact")
	svc = service.NewLoggingService(svc)

	apiServer := api.NewApiServer(svc)
	log.Fatal(apiServer.Start("127.0.0.1:9001"))
}
