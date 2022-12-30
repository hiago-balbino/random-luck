package handler

type requestSchema struct {
	AmountOfGames          uint `form:"amount_of_games"`
	AmountOfNumbersPerGame uint `form:"amount_of_numbers_per_game"`
}
