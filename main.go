package main

import "github.com/hiago-balbino/random-luck/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
