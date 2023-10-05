package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiago-balbino/random-luck/internal/game"
	"github.com/hiago-balbino/random-luck/internal/pkg/apperrors"
	"github.com/hiago-balbino/random-luck/internal/pkg/logger"
)

// web is a struct that implements the Handler interface.
type web struct {
	randomizer game.GameRandomizer
}

// NewWeb is a constructor for creating a new Web handler instance.
func NewWeb(randomize game.GameRandomizer) web {
	return web{randomizer: randomize}
}

// Process is a function implementation to execute calls to create random luck numbers.
func (w web) Process(c *gin.Context) {
	var request requestSchema
	if err := c.BindQuery(&request); err != nil {
		log.Error("error binding query params", logger.FieldError(err))
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": err.Error()})

		return
	}

	games, err := w.randomizer.Randomize(int(request.AmountOfGames), int(request.AmountOfNumbersPerGame))
	if err != nil {
		if errors.Is(err, apperrors.ErrBase) {
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
