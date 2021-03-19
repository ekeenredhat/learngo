// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Enough Picks
//
//  If the player's guess number is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guess number is above 10; then the
//  computer should generate the numbers
//  between 0 and the guess number.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guess number is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
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
	if len(args) < 2 {
		fmt.Println("Pick two numbers")
		return
	}

	// Check if verbose mode flag and set true if so
	vb := false
	if args[0] == "-v" {
		vb = true
	}

	//Convert arguments from String to Int
	guess, err := strconv.Atoi(args[len(args)-2])

	if err != nil {
		fmt.Println("Not a number", err)
		return
	}

	guess2, err2 := strconv.Atoi(args[len(args)-1])
	if err2 != nil {
		fmt.Println("not a number!")
		return
	}

	if (guess <= 0) || (guess2 <= 0) {
		fmt.Println("Pick a positive number")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		var largerNumber int

		if guess > guess2 {
			largerNumber = guess
		} else {
			largerNumber = guess2
		}

		if largerNumber < 10 {
			largerNumber = 10
		}
		n := rand.Intn(largerNumber) + 1
		if vb {
			fmt.Printf("%d ", n)
		}
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
