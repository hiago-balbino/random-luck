package service

import (
	"context"

	"github.com/hiago-balbino/random-luck/internal/core/domain"
)

// GameRandomizer is an interface that handles functions to randomize data to create games.
type GameRandomizer interface {
	// Randomize is an function for randomizing luck numbers to create games.
	Randomize(ctx context.Context, amountOfGames, amountOfNumbersPerGame int) ([]domain.Game, error)
}
