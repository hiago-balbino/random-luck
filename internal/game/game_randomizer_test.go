package game

import (
	"testing"

	"github.com/hiago-balbino/random-luck/v2/internal/pkg/apperrors"
	"github.com/stretchr/testify/assert"
)

func TestRandomize(t *testing.T) {
	testCases := []struct {
		name string
		run  func(*testing.T)
	}{
		{
			name: "should return an error when the amount of games is less than the minimum allowed",
			run: func(t *testing.T) {
				amountOfGames := 0
				amountOfNumbersPerGame := 6

				randomizer := NewGameRandomizer()
				games, err := randomizer.Randomize(amountOfGames, amountOfNumbersPerGame)

				assert.Empty(t, games)
				assert.EqualError(t, err, apperrors.ErrMinAmountOfGames.Error())
			},
		},
		{
			name: "should return an error when the amount of numbers per game is less than the minimum allowed",
			run: func(t *testing.T) {
				amountOfGames := 1
				amountOfNumbersPerGame := 2

				randomizer := NewGameRandomizer()
				games, err := randomizer.Randomize(amountOfGames, amountOfNumbersPerGame)

				assert.Empty(t, games)
				assert.EqualError(t, err, apperrors.ErrMinAmountOfNumbersPerGame.Error())
			},
		},
		{
			name: "should return an error when the amount of numbers per game is greater than the maximum allowed",
			run: func(t *testing.T) {
				amountOfGames := 1
				amountOfNumbersPerGame := 10

				randomizer := NewGameRandomizer()
				games, err := randomizer.Randomize(amountOfGames, amountOfNumbersPerGame)

				assert.Empty(t, games)
				assert.EqualError(t, err, apperrors.ErrMaxAmountOfNumbersPerGame.Error())
			},
		},
		{
			name: "should only generate a game with random numbers among those allowed",
			run: func(t *testing.T) {
				amountOfGames := 1
				amountOfNumbersPerGame := 6

				randomizer := NewGameRandomizer()
				games, err := randomizer.Randomize(amountOfGames, amountOfNumbersPerGame)

				assert.NoError(t, err)
				assert.Equal(t, 1, len(games))
				assertGameNumbers(t, games)
			},
		},
		{
			name: "should generate five games with random numbers among those allowed",
			run: func(t *testing.T) {
				amountOfGames := 5
				amountOfNumbersPerGame := 9

				randomizer := NewGameRandomizer()
				games, err := randomizer.Randomize(amountOfGames, amountOfNumbersPerGame)

				assert.NoError(t, err)
				assert.Equal(t, 5, len(games))
				assertGameNumbers(t, games)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, testCase.run)
	}
}

func assertGameNumbers(t *testing.T, games []Game) {
	for _, game := range games {
		for _, number := range game.Numbers {
			if number < minNumberPerGame || number > maxNumberPerGame {
				t.Fatalf("generated numbers must be between %d and %d", minNumberPerGame, maxNumberPerGame)
			}
		}
	}
}
