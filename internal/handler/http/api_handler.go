package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	errcore "github.com/hiago-balbino/random-luck/internal/core/errors"
	"github.com/hiago-balbino/random-luck/internal/core/service"
	"go.uber.org/zap/zapcore"
)

// APIHandler is a struct that implements the Handler interface.
type APIHandler struct {
	randomizer service.GameRandomizer
}

// NewAPIHandler is a constructor for creating a new APIHandler instance.
func NewAPIHandler(randomizer service.GameRandomizer) APIHandler {
	return APIHandler{randomizer: randomizer}
}

// Process is a function implementation to execute calls to create random luck numbers.
func (h APIHandler) Process(c *gin.Context) {
	var request requestSchema
	if err := c.BindQuery(&request); err != nil {
		log.Error("error binding query params", zapcore.Field{Type: zapcore.StringType, String: err.Error()})
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	games, err := h.randomizer.Randomize(c.Request.Context(), int(request.AmountOfGames), int(request.AmountOfNumbersPerGame))
	if err != nil {
		if errors.Is(err, errcore.ErrBase) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	response := make([]responseSchema, 0)
	for _, game := range games {
		response = append(response, responseSchema{GameID: game.ID, Numbers: game.Numbers})
	}

	c.JSON(http.StatusOK, response)
}
