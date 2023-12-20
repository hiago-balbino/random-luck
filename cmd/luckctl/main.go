package main

import (
	"flag"

	"github.com/hiago-balbino/random-luck/v2/internal/game"
	"github.com/hiago-balbino/random-luck/v2/internal/handler/ctl"
)

func main() {
	amountOfGames := flag.Int("games", 0, "The amount of games")
	amountOfNumbersPerGame := flag.Int("numbers", 0, "The amount of numbers per game")
	flag.Parse()

	randomizer := game.NewGameRandomizer()
	ctl.ProcessRandomLuckNumbers(randomizer, *amountOfGames, *amountOfNumbersPerGame)
}
