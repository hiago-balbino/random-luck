package game

import (
	"context"
	"crypto/rand"
	"math/big"
	"sort"

	"github.com/hiago-balbino/random-luck/internal/pkg/apperrors"
	"github.com/hiago-balbino/random-luck/internal/pkg/logger"
)

// GameRandomizer is an interface that handles functions to randomize data to create games.
type GameRandomizer interface {
	// Randomize is an function for randomizing luck numbers to create games.
	Randomize(ctx context.Context, amountOfGames, amountOfNumbersPerGame int) ([]Game, error)
}

const (
	minAmountOfGamesAllowed   = 1
	minAmountOfNumbersPerGame = 6
	maxAmountOfNumbersPerGame = 9
	minNumberPerGame          = 1
	maxNumberPerGame          = 60
)

var log = logger.GetLogger()

// GameRandomize is a struct that implements GameRandomizer interface that handles functions to randomize data to create games.
type GameRandomize struct{}

// NewGameRandomizer is a constructor for creating a new instance of GameRandomize.
func NewGameRandomizer() GameRandomize {
	return GameRandomize{}
}

// Randomize is an function for randomizing luck numbers to create games.
func (g GameRandomize) Randomize(_ context.Context, amountOfGames, amountOfNumbersPerGame int) ([]Game, error) {
	if err := g.validateParameters(amountOfGames, amountOfNumbersPerGame); err != nil {
		log.Error("validate parameters", logger.FieldError(err))

		return nil, err
	}

	games := make([]Game, 0)
	for i := minAmountOfGamesAllowed; i <= amountOfGames; i++ {
		numbers := make([]int, 0)

		for j := minNumberPerGame; j <= amountOfNumbersPerGame; j++ {
			number := g.generateNewNumber(numbers)
			numbers = append(numbers, number)
		}

		sort.Ints(numbers)
		games = append(games, Game{ID: i, Numbers: numbers})
	}

	return games, nil
}

// validateParameters checks if the parameters are in good condition to process the numbers.
func (g GameRandomize) validateParameters(amountOfGames, amountOfNumbersPerGame int) error {
	switch {
	case amountOfGames < minAmountOfGamesAllowed:
		return apperrors.ErrMinAmountOfGames
	case amountOfNumbersPerGame < minAmountOfNumbersPerGame:
		return apperrors.ErrMinAmountOfNumbersPerGame
	case amountOfNumbersPerGame > maxAmountOfNumbersPerGame:
		return apperrors.ErrMaxAmountOfNumbersPerGame
	default:
		return nil
	}
}

// generateNewNumber is used to create new positive number between 1 and 60.
func (g GameRandomize) generateNewNumber(numbersAlreadyGenerated []int) int {
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(maxNumberPerGame-minNumberPerGame))
	if err != nil {
		log.Error("generate new randomNumber", logger.FieldError(err))

		return g.generateNewNumber(numbersAlreadyGenerated)
	}
	newNumber := int(randomNumber.Add(randomNumber, big.NewInt(minNumberPerGame)).Int64())

	for _, numberAlreadyGenerated := range numbersAlreadyGenerated {
		if numberAlreadyGenerated == newNumber {
			return g.generateNewNumber(numbersAlreadyGenerated)
		}
	}

	return newNumber
}
