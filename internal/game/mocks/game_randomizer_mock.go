package mocks

import (
	"github.com/hiago-balbino/random-luck/v2/internal/game"
	"github.com/stretchr/testify/mock"
)

// GameRandomizerMock is a mock struct that implements the GameRandomizer interface.
type GameRandomizerMock struct {
	mock.Mock
}

// Randomize is an function for randomizing luck numbers to create games.
func (m *GameRandomizerMock) Randomize(amountOfGames, amountOfNumbersPerGame int) ([]game.Game, error) {
	args := m.Called(amountOfGames, amountOfNumbersPerGame)

	return args.Get(0).([]game.Game), args.Error(1)
}
