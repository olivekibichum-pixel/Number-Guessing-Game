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
		var difference = (guess - secretNumber)

		fmt.Print("Guess a number: ")
		fmt.Scan(&guess)

		if difference < 0 {
			difference = -difference
		}

		if difference <= 3 {
			fmt.Println("Warmer!")
		} else if guess == secretNumber {
			fmt.Println("Correct!")
			break
		} else if guess > secretNumber {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Too low!")
		}
	}
}
