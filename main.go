package main

import (
	"fmt"
	"math/rand"
	"sort"
)

const (
	minNumberToGame = 1
	maxNumberToGame = 61
)

func main() {
	fmt.Println("** Please enter the numbers to generate the games **")
	numberOfGames, numbersToGame := readInput()
	fmt.Println("")

	fmt.Println("** Generating **")
	result := generateNumbers(numberOfGames, numbersToGame)
	fmt.Println("Total games generated:", len(result))
	fmt.Println("")

	fmt.Println("** Results: **")
	printOutput(result)
}

func readInput() (int, int) {
	var numberOfGames int
	var numbersToGame int

	fmt.Println("Select number of games:")
	_, err := fmt.Scanln(&numberOfGames)
	if err != nil {
		panic("error while scanning the number of games")
	}

	fmt.Println("Select numbers to game:")
	_, err = fmt.Scanln(&numbersToGame)
	if err != nil {
		panic("error while scanning the numbers to game")
	}
	return numberOfGames, numbersToGame
}

func generateNumbers(numberOfGames int, numbersToGame int) map[int][]int {
	result := make(map[int][]int)
	for i := minNumberToGame; i <= numberOfGames; i++ {
		var numbers []int
		for j := minNumberToGame; j <= numbersToGame; j++ {
			numbers = append(numbers, generate())
		}
		result[i] = numbers
	}
	return result
}

func generate() int {
	value := rand.Intn(maxNumberToGame)
	if value < minNumberToGame {
		return generate()
	}
	return value
}

func printOutput(result map[int][]int) {
	keys := sortKeys(result)

	for _, key := range keys {
		fmt.Println("Numbers to game:", key)
		fmt.Println(result[key])
		fmt.Println("")
	}
}

func sortKeys(result map[int][]int) []int {
	var keys []int
	for key := range result {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
}
