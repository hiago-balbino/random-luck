package errors

import "errors"

var (
	ErrMinAmountOfGames          = errors.New("amount of games is less than the minimum allowed (amountOfGames>=1)")
	ErrMinAmountOfNumbersPerGame = errors.New("amount of numbers per game is less than the minimum allowed (amountOfNumbersPerGame>=6)")
	ErrMaxAmountOfNumbersPerGame = errors.New("amount of numbers per game is greater than the maximum allowed (amountOfNumbersPerGame<=9)")
)
