package main

import (
	"log"

	news "github.com/Le0nar/crud_go"
	"github.com/Le0nar/crud_go/pkg/handler"
	"github.com/Le0nar/crud_go/pkg/repository"
	"github.com/Le0nar/crud_go/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(news.Server)
	port := viper.GetString("port")
	if err := srv.Run(port, handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
