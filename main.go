package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")

	var secretNumber = rand.Intn(100) + 1

	for {
		var guess int
		fmt.Print("Guess a number: ")
		fmt.Scan(&guess)

		if guess == secretNumber {
			fmt.Println("Correct!")
			break //stops the loop when tghe player gets the correct answer.
		} else if guess > secretNumber {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Too low!")
		}
	}
}
