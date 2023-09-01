package main

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
)

func main() {
	var toGuess int
	var nonNumericRegex = regexp.MustCompile(`[^0-9 ]+`)
	guesses := 0
	guessed := false

	// Sees the number
	toGuess = rand.Intn(99) + 1 // To guess between 1-100
	fmt.Println("Guess the number...")
	fmt.Println("Between 1 and 100")

	for !guessed {
		var guess string
		var guessInt int

		fmt.Println("Enter your guess:")
		_, err := fmt.Scanln(&guess)
		if err != nil {
			log.Fatal(err)
		}

		guess = nonNumericRegex.ReplaceAllString(guess, "")

		guessInt, err = strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}

		if toGuess == guessInt {
			fmt.Println("Congratulations! You guessed the number correctly!")
			fmt.Printf("It took %d guesses!\n", guesses)
			break
		}

		if guessInt < toGuess {
			fmt.Println("Higher")
		}
		if guessInt > toGuess {
			fmt.Println("lower")
		}

		guesses++

	}
}
