package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	errcore "github.com/hiago-balbino/random-luck/internal/core/errors"
	"github.com/hiago-balbino/random-luck/internal/core/service"
	"go.uber.org/zap/zapcore"
)

// WebHandler is a struct that implements the Handler interface.
type WebHandler struct {
	randomizer service.GameRandomizer
}

// NewWebHandler is a constructor for creating a new WebHandler instance.
func NewWebHandler(randomize service.GameRandomizer) WebHandler {
	return WebHandler{randomizer: randomize}
}

// Process is a function implementation to execute calls to create random luck numbers.
func (h WebHandler) Process(c *gin.Context) {
	var request requestSchema
	if err := c.BindQuery(&request); err != nil {
		log.Error("error binding query params", zapcore.Field{Type: zapcore.StringType, String: err.Error()})
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": err.Error()})

		return
	}

	games, err := h.randomizer.Randomize(c.Request.Context(), int(request.AmountOfGames), int(request.AmountOfNumbersPerGame))
	if err != nil {
		if errors.Is(err, errcore.ErrBase) {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": err.Error()})

			return
		}
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"message": err.Error()})

		return
	}

	response := make([]responseSchema, 0)
	for _, game := range games {
		response = append(response, responseSchema{GameID: game.ID, Numbers: game.Numbers})
	}

	c.HTML(http.StatusOK, "games.html", gin.H{"games": response})
}
