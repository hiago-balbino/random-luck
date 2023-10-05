package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/hiago-balbino/random-luck/internal/game"
	"github.com/hiago-balbino/random-luck/internal/game/mocks"
)

func TestProcess(t *testing.T) {
	amountOfGames := 2
	amountOfNumbersPerGame := 6
	randomizer := game.NewGameRandomizer()

	t.Run("web handler", func(t *testing.T) {
		handler := setupHandler(randomizer)
		server := httptest.NewServer(handler)
		defer server.Close()

		e := httpexpect.Default(t, server.URL)

		t.Run("should return 4xx error", func(t *testing.T) {
			t.Run("when negative amount of games query param", func(t *testing.T) {
				e.GET("/process").
					WithQuery("amount_of_games", -1).
					WithQuery("amount_of_numbers_per_game", amountOfNumbersPerGame).
					Expect().
					Status(http.StatusBadRequest).
					Body().
					Contains("strconv.ParseUint: parsing").
					Contains("invalid syntax")
			})
			t.Run("when zero amount of games query param", func(t *testing.T) {
				e.GET("/process").
					WithQuery("amount_of_games", 0).
					WithQuery("amount_of_numbers_per_game", amountOfNumbersPerGame).
					Expect().
					Status(http.StatusBadRequest).
					Body().Contains("amount of games is less than the minimum allowed")
			})
			t.Run("when negative amount of numbers per game query param", func(t *testing.T) {
				e.GET("/process").
					WithQuery("amount_of_games", amountOfGames).
					WithQuery("amount_of_numbers_per_game", -1).
					Expect().
					Status(http.StatusBadRequest).
					Body().
					Contains("strconv.ParseUint: parsing").
					Contains("invalid syntax")
			})
			t.Run("when zero amount of numbers per game query param", func(t *testing.T) {
				e.GET("/process").
					WithQuery("amount_of_games", amountOfGames).
					WithQuery("amount_of_numbers_per_game", 0).
					Expect().
					Status(http.StatusBadRequest).
					Body().Contains("amount of numbers per game is less than the minimum allowed")
			})
			t.Run("when exceeded amount of numbers per game query param", func(t *testing.T) {
				e.GET("/process").
					WithQuery("amount_of_games", amountOfGames).
					WithQuery("amount_of_numbers_per_game", 100).
					Expect().
					Status(http.StatusBadRequest).
					Body().Contains("amount of numbers per game is greater than the maximum allowed")
			})
		})

		t.Run("should return 2xx", func(t *testing.T) {
			t.Run("when to successfully generate random luck numbers", func(t *testing.T) {
				e.GET("/process").
					WithQuery("amount_of_games", amountOfGames).
					WithQuery("amount_of_numbers_per_game", amountOfNumbersPerGame).
					Expect().
					Status(http.StatusOK).
					Body().
					Contains("Game 1").
					Contains("Game 2")
			})
		})

		t.Run("should return 5xx error", func(t *testing.T) {
			unexpectedErr := errors.New("unexpected error")
			randomizer := new(mocks.GameRandomizerMock)
			randomizer.On("Randomize", amountOfGames, amountOfNumbersPerGame).Return([]game.Game{}, unexpectedErr)

			handler := setupHandler(randomizer)
			server := httptest.NewServer(handler)
			defer server.Close()

			e := httpexpect.Default(t, server.URL)
			e.GET("/process").
				WithQuery("amount_of_games", amountOfGames).
				WithQuery("amount_of_numbers_per_game", amountOfNumbersPerGame).
				Expect().
				Status(http.StatusInternalServerError).
				Body().
				Contains(unexpectedErr.Error())
		})
	})
}

func setupHandler(randomizer game.GameRandomizer) *gin.Engine {
	handler := NewWeb(randomizer)
	server := Server{handler: handler}
	router := server.setupRoutes("../../../web/templates/*")

	return router
}
