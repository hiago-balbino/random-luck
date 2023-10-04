package ctl

import (
	"context"

	"github.com/fatih/color"
	"github.com/hiago-balbino/random-luck/internal/game"
	"github.com/rodaine/table"
)

// ProcessRandomLuckNumbers generates random luck numbers for a given amount of games
// using the provided randomizer and outputs the generated games.
func ProcessRandomLuckNumbers(randomizer game.GameRandomizer, amountOfGames, amountOfNumbersPerGame int) {
	if games, _ := randomizer.Randomize(context.Background(), amountOfGames, amountOfNumbersPerGame); len(games) > 0 {
		outputAsTable(games)
	}
}

func outputAsTable(games []game.Game) {
	tbl := table.New("Game", "Numbers").
		WithHeaderFormatter(color.New(color.FgCyan, color.Bold).SprintfFunc()).
		WithFirstColumnFormatter(color.New(color.FgCyan).SprintfFunc())

	for _, game := range games {
		tbl.AddRow(game.ID, game.Numbers)
	}

	tbl.Print()
}
