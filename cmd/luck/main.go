package main

import (
	"github.com/hiago-balbino/random-luck/config"
	handler "github.com/hiago-balbino/random-luck/internal/handler/http"
)

func main() {
	config.InitConfigurations()
	server := handler.NewServer(handler.API)
	server.Start()
}
