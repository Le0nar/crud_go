package main

import (
	"log"

	news "github.com/Le0nar/crud_go"
	"github.com/Le0nar/crud_go/pkg/handler"
	"github.com/Le0nar/crud_go/pkg/repository"
	"github.com/Le0nar/crud_go/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(news.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}
