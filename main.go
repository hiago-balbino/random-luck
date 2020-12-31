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
			number := generate()
			if isAlreadyAdded(number, numbers) {
				j--
				continue
			}
			numbers = append(numbers, number)
		}
		sort.Ints(numbers)
		result[i] = numbers
	}
	return result
}

func generate() int {
	number := rand.Intn(maxNumberToGame)
	if number < minNumberToGame {
		return generate()
	}
	return number
}

func isAlreadyAdded(number int, numbers []int) bool {
	for _, n := range numbers {
		if n == number {
			return true
		}
	}
	return false
}

func printOutput(result map[int][]int) {
	for i := 1; i <= len(result); i++ {
		fmt.Println("Numbers to game:", i)
		fmt.Println(result[i])
		fmt.Println("")
	}
}
