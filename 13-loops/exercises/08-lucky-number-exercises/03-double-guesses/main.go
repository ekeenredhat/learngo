// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus
// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//declare constants
const (
	maxTurns = 5
	usage    = `Welcome message. THis program will be %d random numbers.
	Good luck asshole!`
)

func main() {

	//generate randon number
	rand.Seed(time.Now().UnixNano())

	//Capture arguments, ignoring the first point and capturing to the end of the slice
	args := os.Args[1:]

	//Ensure that they provide two numbers
	if len(args) != 2 {
		fmt.Println("Pick two numbers")
		return
	}

	// A few options here;
	// 1. Create a for loop to iterate through the len(args) to create Guess1 and Guess2 Ints
	// 2. Create a separate method to validate and return acceptable Guess1 and Guess2
	// 3. Create a slice for the Guesses
	// 4. Simple approach, create a Guess2 without the loop

	//Convert arguments from String to Int
	guess, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("Not a number", err)
		return
	}

	guess2, err2 := strconv.Atoi(args[1])
	if err2 != nil {
		fmt.Println("not a number!")
		return
	}

	if (guess <= 0) || (guess > 10) || (guess2 <= 0) || (guess2 > 10) {
		fmt.Println("Pick a positive number, between 1 and 10")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		var largerNumber int

		if guess > guess2 {
			largerNumber = guess
		} else {
			largerNumber = guess2
		}

		n := rand.Intn(largerNumber) + 1

		if n == guess || n == guess2 {
			switch rand.Intn(3) {
			case 0:
				fmt.Println("You win, mate!")
			case 1:
				fmt.Println("Have a nice day, winner")
			case 2:
				fmt.Println("Great job winning!")
			}
			return
		}
	}
	msg := "%s try again?\n"

	switch rand.Intn(2) {
	case 0:
		fmt.Printf(msg, "* You lost")
	case 1:
		fmt.Printf(msg, "* Oh no, loser")

	}
}
