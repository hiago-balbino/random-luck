package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	coreservice "github.com/hiago-balbino/random-luck/internal/core/service"
	"github.com/hiago-balbino/random-luck/internal/pkg/service"
)

func TestProcess_WebAndAPI(t *testing.T) {
	givenAmountOfGames := 2
	givenAmountOfNumbersPerGame := 6
	randomizer := service.NewGameRandomizer()

	t.Run("web handler", func(t *testing.T) {
		handler := setupHandler(WEB, randomizer)
		server := httptest.NewServer(handler)
		defer server.Close()

		e := httpexpect.Default(t, server.URL)

		t.Run("should return 4xx error", func(t *testing.T) {
			t.Run("when negative amount of games query param", func(t *testing.T) {
				e.GET("/process").
					WithQuery("amount_of_games", -1).
					WithQuery("amount_of_numbers_per_game", givenAmountOfNumbersPerGame).
					Expect().
					Status(http.StatusBadRequest).
					Body().
					Contains("strconv.ParseUint: parsing").
					Contains("invalid syntax")
			})
			t.Run("when zero amount of games query param", func(t *testing.T) {
				e.GET("/process").
					WithQuery("amount_of_games", 0).
					WithQuery("amount_of_numbers_per_game", givenAmountOfNumbersPerGame).
					Expect().
					Status(http.StatusBadRequest).
					Body().Contains("amount of games is less than the minimum allowed")
			})
			t.Run("when negative amount of numbers per game query param", func(t *testing.T) {
				e.GET("/process").
					WithQuery("amount_of_games", givenAmountOfGames).
					WithQuery("amount_of_numbers_per_game", -1).
					Expect().
					Status(http.StatusBadRequest).
					Body().
					Contains("strconv.ParseUint: parsing").
					Contains("invalid syntax")
			})
			t.Run("when zero amount of numbers per game query param", func(t *testing.T) {
				e.GET("/process").
					WithQuery("amount_of_games", givenAmountOfGames).
					WithQuery("amount_of_numbers_per_game", 0).
					Expect().
					Status(http.StatusBadRequest).
					Body().Contains("amount of numbers per game is less than the minimum allowed")
			})
			t.Run("when exceeded amount of numbers per game query param", func(t *testing.T) {
				e.GET("/process").
					WithQuery("amount_of_games", givenAmountOfGames).
					WithQuery("amount_of_numbers_per_game", 100).
					Expect().
					Status(http.StatusBadRequest).
					Body().Contains("amount of numbers per game is greater than the maximum allowed")
			})
		})

		t.Run("should return 2xx", func(t *testing.T) {
			t.Run("when to successfully generate random luck numbers", func(t *testing.T) {
				e.GET("/process").
					WithQuery("amount_of_games", givenAmountOfGames).
					WithQuery("amount_of_numbers_per_game", givenAmountOfNumbersPerGame).
					Expect().
					Status(http.StatusOK).
					Body().
					Contains("Game 1").
					Contains("Game 2")
			})
		})
	})

	t.Run("API handler", func(t *testing.T) {
		handler := setupHandler(API, randomizer)
		server := httptest.NewServer(handler)
		defer server.Close()

		e := httpexpect.Default(t, server.URL)

		t.Run("should return 4xx error", func(t *testing.T) {
			t.Run("when negative amount of games query param", func(t *testing.T) {
				e.GET("/process").
					WithQuery("amount_of_games", -1).
					WithQuery("amount_of_numbers_per_game", givenAmountOfNumbersPerGame).
					Expect().
					Status(http.StatusBadRequest).
					Body().Contains(`{"message":"strconv.ParseUint: parsing \"-1\": invalid syntax"}`)
			})
			t.Run("when zero amount of games query param", func(t *testing.T) {
				e.GET("/process").
					WithQuery("amount_of_games", 0).
					WithQuery("amount_of_numbers_per_game", givenAmountOfNumbersPerGame).
					Expect().
					Status(http.StatusBadRequest).
					Body().Contains("amount of games is less than the minimum allowed")
			})
			t.Run("when negative amount of numbers per game query param", func(t *testing.T) {
				e.GET("/process").
					WithQuery("amount_of_games", givenAmountOfGames).
					WithQuery("amount_of_numbers_per_game", -1).
					Expect().
					Status(http.StatusBadRequest).
					Body().Contains(`{"message":"strconv.ParseUint: parsing \"-1\": invalid syntax"}`)
			})
			t.Run("when zero amount of numbers per game query param", func(t *testing.T) {
				e.GET("/process").
					WithQuery("amount_of_games", givenAmountOfGames).
					WithQuery("amount_of_numbers_per_game", 0).
					Expect().
					Status(http.StatusBadRequest).
					Body().Contains("amount of numbers per game is less than the minimum allowed")
			})
			t.Run("when exceeded amount of numbers per game query param", func(t *testing.T) {
				e.GET("/process").
					WithQuery("amount_of_games", givenAmountOfGames).
					WithQuery("amount_of_numbers_per_game", 100).
					Expect().
					Status(http.StatusBadRequest).
					Body().Contains("amount of numbers per game is greater than the maximum allowed")
			})
		})

		t.Run("should return 2xx", func(t *testing.T) {
			t.Run("when to successfully generate random luck numbers", func(t *testing.T) {
				elements := e.GET("/process").
					WithQuery("amount_of_games", givenAmountOfGames).
					WithQuery("amount_of_numbers_per_game", givenAmountOfNumbersPerGame).
					Expect().
					Status(http.StatusOK).
					JSON().Array()

				elements.Length().Equal(givenAmountOfGames)

				elements.Element(0).
					Object().
					ValueEqual("game_id", 1).
					Value("numbers").Array().Length().Equal(6)

				elements.Element(1).
					Object().
					ValueEqual("game_id", 2).
					Value("numbers").Array().Length().Equal(6)
			})
		})
	})
}

func setupHandler(serverType int, randomizer coreservice.GameRandomizer) *gin.Engine {
	if serverType == WEB {
		handler := NewWebHandler(randomizer)
		server := Server{serverType: serverType, handler: handler}
		router := server.setupRoutes("../templates/*")

		return router
	}

	handler := NewAPIHandler(randomizer)
	server := Server{serverType: serverType, handler: handler}
	router := server.setupRoutes("")

	return router
}
