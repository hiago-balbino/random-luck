package handler

import (
	"fmt"
	"net/http"

	"github.com/arl/statsviz"
	"github.com/gin-gonic/gin"
	"github.com/hiago-balbino/random-luck/config"
	"github.com/hiago-balbino/random-luck/internal/game"
	"github.com/hiago-balbino/random-luck/internal/pkg/logger"
	"github.com/spf13/viper"
)

var log = logger.GetLogger()

// Server is an structure to support the handler.
type Server struct {
	handler Handler
}

// NewServer create a new instante of Server structure.
func NewServer() Server {
	config.InitConfigurations()
	randomizer := game.NewGameRandomizer()

	return Server{
		handler: NewWeb(randomizer),
	}
}

// Start initialize the router.
func (s Server) Start() {
	router := s.setupRoutes("web/templates/*")

	if err := router.Run(fmt.Sprintf(":%s", viper.GetString("PORT"))); err != nil {
		log.Fatal("error while server starting", logger.FieldError(err))
	}
}

func (s Server) setupRoutes(templatePath string) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob(templatePath)
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/metrics/*filepath", func(c *gin.Context) {
		if c.Param("filepath") == "/ws" {
			statsviz.Ws(c.Writer, c.Request)

			return
		}
		statsviz.IndexAtRoot("/metrics").ServeHTTP(c.Writer, c.Request)
	})
	router.GET("/process", s.handler.Process)

	return router
}
