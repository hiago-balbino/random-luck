package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hiago-balbino/random-luck/configuration"
	"github.com/hiago-balbino/random-luck/internal/pkg/logger"
	"github.com/hiago-balbino/random-luck/internal/pkg/service"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log = logger.GetLogger()

// Server is an structure to support the handler.
type Server struct {
	serverType int
	handler    Handler
}

// NewServer create a new instante of Server structure.
func NewServer(serverType int) Server {
	configuration.InitConfigurations()
	randomizer := service.NewGameRandomizer()

	if serverType == WEB {
		return Server{handler: NewWebHandler()}
	}

	return Server{
		serverType: serverType,
		handler:    NewAPIHandler(randomizer),
	}
}

// Start initialize the router.
func (s Server) Start() {
	router := s.setupRoutes("templates/*")

	if err := router.Run(fmt.Sprintf(":%s", viper.GetString("PORT"))); err != nil {
		log.Fatal("error while server starting", zap.Field{Type: zapcore.StringType, String: err.Error()})
	}
}

func (s Server) setupRoutes(templatePath string) *gin.Engine {
	router := gin.Default()

	if s.serverType == WEB {
		router.LoadHTMLGlob(templatePath)
		router.GET("/index", func(_ *gin.Context) {})
	}
	router.GET("/process", s.handler.Process)

	return router
}
