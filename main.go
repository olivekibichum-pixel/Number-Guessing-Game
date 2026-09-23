package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to the Number Guessing Game! \nThe number range is 1-100. \nTo EXIT the game, enter 0. \nHave Fun!")

	var secretNumber = rand.Intn(100) + 1

	for {
		var guess int

		fmt.Print("Guess the number: ")
		fmt.Scan(&guess)

		var difference int = (guess - secretNumber)

		if difference < 0 {
			difference = -difference
		}

		if guess == 0 {
			break
		} else if guess > 100 {
			fmt.Println("Error")
		} else if guess == secretNumber {
			fmt.Println("Correct!")
			break
		} else if difference <= 3 {
			fmt.Println("Warmer!")
		} else if guess > secretNumber {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Too low!")
		}
	}
}
