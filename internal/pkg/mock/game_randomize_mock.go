package mock

import (
	"context"

	"github.com/hiago-balbino/random-luck/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

// GameRandomizeMock is an mock of GameRandomizer interface that handles functions to randomize data to create games.
type GameRandomizeMock struct {
	mock.Mock
}

// Randomize is a mock function for randomizing luck numbers to create games.
func (g *GameRandomizeMock) Randomize(ctx context.Context, amountOfGames, amountOfNumbersPerGame int) ([]domain.Game, error) {
	args := g.Called(ctx, amountOfGames, amountOfNumbersPerGame)

	return args.Get(0).([]domain.Game), args.Error(1)
}
