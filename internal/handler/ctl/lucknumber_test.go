package ctl

import (
	"testing"

	"github.com/hiago-balbino/random-luck/internal/game"
	"github.com/hiago-balbino/random-luck/internal/game/mocks"
	"github.com/hiago-balbino/random-luck/internal/pkg/apperrors"
)

func TestProcessRandomLuckNumbers(t *testing.T) {
	testCases := []struct {
		name string
		run  func(*testing.T)
	}{
		{
			name: "should not panic when randomizer returns an error",
			run: func(t *testing.T) {
				amountOfGames := 0
				amountOfNumbersPerGame := 6

				randomizer := new(mocks.GameRandomizerMock)
				randomizer.On("Randomize", amountOfGames, amountOfNumbersPerGame).Return([]game.Game{}, apperrors.ErrMinAmountOfGames)

				ProcessRandomLuckNumbers(randomizer, amountOfGames, amountOfNumbersPerGame)

				randomizer.AssertExpectations(t)
				randomizer.AssertCalled(t, "Randomize", amountOfGames, amountOfNumbersPerGame)
			},
		},
		{
			name: "should process the luck numbers successfully",
			run: func(t *testing.T) {
				amountOfGames := 2
				amountOfNumbersPerGame := 6

				randomizer := new(mocks.GameRandomizerMock)
				randomizer.On("Randomize", amountOfGames, amountOfNumbersPerGame).Return([]game.Game{
					{ID: 1, Numbers: []int{1, 2, 3, 4, 5, 6}},
					{ID: 2, Numbers: []int{1, 2, 3, 4, 5, 6}},
				}, nil)

				ProcessRandomLuckNumbers(randomizer, amountOfGames, amountOfNumbersPerGame)

				randomizer.AssertExpectations(t)
				randomizer.AssertCalled(t, "Randomize", amountOfGames, amountOfNumbersPerGame)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, testCase.run)
	}
}
