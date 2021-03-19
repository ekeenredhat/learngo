// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Dynamic Difficulty
//
//  Current game picks only 5 numbers (5 turns).
//
//  Make sure that the game adjust its own difficulty
//  depending on the guess number.
//
// RESTRICTION
//  Do not make the game to easy. Only adjust the
//  difficulty if the guess is above 10.
//
// EXPECTED OUTPUT
//  Suppose that the player runs the game like this:
//    go run main.go 5
//
//    Then the computer should pick 5 random numbers.
//
//  Or, if the player runs it like this:
//    go run main.go 25
//
//    Then the computer may pick 11 random numbers
//    instead.
//
//  Or, if the player runs it like this:
//    go run main.go 100
//
//    Then the computer may pick 30 random numbers
//    instead.
//
//  As you can see, greater guess number causes the
//  game to increase the game turns, which in turn
//  adjust the game's difficulty dynamically.
//
//  Because, greater guess number makes it harder to win.
//  But, this way, game's difficulty will be dynamic.
//  It will adjust its own difficulty depending on the
//  guess number.
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
	usage = `Welcome message. THis program will be %d random numbers.
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

	var largerNumber int
	var maxTurns int

	if guess > guess2 {
		largerNumber = guess
	} else {
		largerNumber = guess2
	}

	if largerNumber > 15 && largerNumber < 30 {
		maxTurns = largerNumber / 2
	}
	if largerNumber >= 30 && largerNumber < 50 {
		maxTurns = largerNumber / 4
	}
	if largerNumber >= 50 && largerNumber < 100 {
		maxTurns = largerNumber / 6
	}

	for turn := 0; turn < maxTurns; turn++ {
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
