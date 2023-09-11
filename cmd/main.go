package main

import (
	"log"

	news "github.com/Le0nar/crud_go"
	"github.com/Le0nar/crud_go/pkg/handler"
)

func main() {
	handler := new(handler.Handler)

	srv := new(news.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}