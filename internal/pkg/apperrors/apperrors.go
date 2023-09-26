package apperrors

import (
	"errors"
	"fmt"
)

var (
	// ErrBase is an error that marks other errors as known by the application.
	ErrBase = errors.New("")
	// ErrMinAmountOfGames is an error that represents that the amount of games is less than the minimum allowed.
	ErrMinAmountOfGames = fmt.Errorf("amount of games is less than the minimum allowed (amountOfGames>=1)%w", ErrBase)
	// ErrMinAmountOfNumbersPerGame is an error that represents that the amount of numbers per game is less than the minimum allowed.
	ErrMinAmountOfNumbersPerGame = fmt.Errorf("amount of numbers per game is less than the minimum allowed (amountOfNumbersPerGame>=6)%w", ErrBase)
	// ErrMaxAmountOfNumbersPerGame is an error that represents that the amount of numbers per game is greater than the maximum allowed.
	ErrMaxAmountOfNumbersPerGame = fmt.Errorf("amount of numbers per game is greater than the maximum allowed (amountOfNumbersPerGame<=9)%w", ErrBase)
)
