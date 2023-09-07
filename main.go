package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
)

type game struct {
	gameOver bool
	toGuess  int
	guess    int
	guesses  int
}

func main() {
	g := game{}
	g.init()
	g.titleText()
	g.loop()
}

// init sets the initial state of the game.
func (g *game) init() {
	// Seed the number
	g.toGuess = rand.Intn(99) + 1 // To guess between 1-100
	g.gameOver = false
	g.guesses = 0
}

// loop the main game loop where the guessing happens
func (g *game) loop() {
	for !g.gameOver {

		fmt.Println("Enter your guess:")
		guess, err := g.getInput()
		if err != nil {
			fmt.Println("Invalid guess, please enter a number from 1 to 100")
			continue
		}

		g.guesses++

		if g.toGuess == guess {
			fmt.Println("Congratulations! You guessed the number correctly!")
			fmt.Printf("It took %d guesses!\n", g.guesses)
			g.gameOver = true
		}

		if guess < g.toGuess {
			fmt.Println("Higher")
		}
		if guess > g.toGuess {
			fmt.Println("Lower")
		}
	}
	g.playAgain()
}

// playAgain on completion of a successful game the user is asked if they would like to play again.
func (g *game) playAgain() {
	fmt.Println("Would you like to play again? (Y/N)")
	var playAgain string
	_, err := fmt.Scanln(&playAgain)
	if err != nil {
		return
	}

	if playAgain == "y" || playAgain == "Y" {
		g.init()
		g.titleText()
		g.loop()

	}
}

// getInput asks the user for the number.
// Sanitizes any erroneous input returning just the guess or err
func (g *game) getInput() (guess int, err error) {
	var nonNumericRegex = regexp.MustCompile(`[^0-9 ]+`)
	var guessString string
	_, err = fmt.Scanln(&guessString)
	if err != nil {
		return 0, err
	}

	guessString = nonNumericRegex.ReplaceAllString(guessString, "")

	guess, err = strconv.Atoi(guessString)
	if err != nil {
		return 0, err
	}

	return guess, nil

}

func (g *game) titleText() {
	fmt.Println("Guess the number...")
	fmt.Println("Between 1 and 100")
}
