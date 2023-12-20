package main

import (
	"github.com/hiago-balbino/random-luck/v2/config"
	handler "github.com/hiago-balbino/random-luck/v2/internal/handler/http"
)

func main() {
	config.InitConfigurations()
	server := handler.NewServer()
	server.Start()
}
