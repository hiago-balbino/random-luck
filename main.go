package main

import (
	"fmt"
	"math/rand"
	"sort"
)

const (
	minNumberToGame   = 1
	maxNumberToGame   = 61
	minNumbersAllowed = 6
	maxNumbersAllowed = 9
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

// readInput read the number of games and numbers to game validating the entries
func readInput() (int, int) {
	var numberOfGames int
	var numbersToGame int

	fmt.Println("Select number of games:")
	_, err := fmt.Scanln(&numberOfGames)
	if err != nil {
		panic("error while scanning the number of games")
	}
	if numberOfGames < minNumberToGame {
		panic(fmt.Sprintf("the minimum number of games is %d", minNumberToGame))
	}

	fmt.Println("Select numbers to game:")
	_, err = fmt.Scanln(&numbersToGame)
	if err != nil {
		panic("error while scanning the numbers to game")
	}
	if numbersToGame < minNumbersAllowed || numbersToGame > maxNumbersAllowed {
		panic(fmt.Sprintf("the range numbers allowed to game is between %d and %d", minNumbersAllowed, maxNumbersAllowed))
	}
	return numberOfGames, numbersToGame
}

// generateNumbers is used to create numbers by entries and returns ordered result
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

// generate is used to create positive numbers between 1 and 60
func generate() int {
	number := rand.Intn(maxNumberToGame)
	if number < minNumberToGame {
		return generate()
	}
	return number
}

// isAlreadyAdded check if the number already exists in the result collection
func isAlreadyAdded(number int, numbers []int) bool {
	for _, n := range numbers {
		if n == number {
			return true
		}
	}
	return false
}

// printOutput print in the output all the numbers generated
func printOutput(result map[int][]int) {
	for i := 1; i <= len(result); i++ {
		fmt.Println("Numbers to game:", i)
		fmt.Println(result[i])
		fmt.Println("")
	}
}
