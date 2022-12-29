package service

import (
	"context"

	"github.com/hiago-balbino/random-luck/internal/core/domain"
)

// GameRandomize is a struct that implements GameRandomizer interface that handles functions to randomize data to create games.
type GameRandomize struct{}

// NewGameRandomize is a constructor for creating a new instance of GameRandomize.
func NewGameRandomize() GameRandomize {
	return GameRandomize{}
}

// Randomize is an function for randomizing luck numbers to create games.
func (g GameRandomize) Randomize(ctx context.Context, amountOfGames, amountNumbersPerGame int) ([]domain.Game, error) {
	return nil, nil
}
