package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to the Number Guessing Game! \nChoose Game Level \n 1.Easy \n 2.Medium \n 3.Hard")

	var level int
	fmt.Println("Enter level: ")
	fmt.Scan(&level)

	var minNumber int
	var maxNumber int
	var warmerDifference int

	if level == 1 {
		minNumber = 1
		maxNumber = 100
		warmerDifference = 3
	} else if level == 2 {
		minNumber = 101
		maxNumber = 999
		warmerDifference = 10
	} else if level == 3 {
		minNumber = 1000
		maxNumber = 9999
		warmerDifference = 20
	} else {
		fmt.Println("Invalid level!")
	}

	var secretNumber int = rand.Intn(maxNumber-minNumber+1) - minNumber

	fmt.Println("The number range is ", minNumber, "-", maxNumber)
	fmt.Println("To EXIT the game, enter 0 \nHave Fun!")

gameLoop:
	for {
		var guess int

		fmt.Print("Guess the number: ")
		fmt.Scan(&guess)

		var difference int = (guess - secretNumber)

		if difference < 0 {
			difference = -difference
		}

		switch {
		case guess == 0:
			fmt.Println("Exiting...")
			break gameLoop
		case guess < minNumber || guess > maxNumber:
			fmt.Println("Error")
		case guess == secretNumber:
			fmt.Println("Correct!")
			break gameLoop
		case difference <= warmerDifference:
			fmt.Println("Warmer!")
		case guess > secretNumber:
			fmt.Println("Too high!")
		default:
			fmt.Println("Too low!")
		}

	}
}
