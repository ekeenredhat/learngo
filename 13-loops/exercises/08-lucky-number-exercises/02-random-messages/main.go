// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

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

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.

Wanna play?
`
)

func main() {

	//generate randon number
	rand.Seed(time.Now().UnixNano())

	//Capture arguments, ignoring the first point and capturing to the end of the slice
	args := os.Args[1:]

	//Ensure that the arguments isn't empty
	if len(args) != 1 {
		fmt.Println("Pick a number")
		return
	}
	//Convert arguments from String to Int
	guess, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("Not a number", err)
		return
	}

	if (guess <= 0) || (guess > 10) {
		fmt.Println("Pick a positive number, between 1 and 10")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess) + 1

		if n == guess {
			switch rand.Intn(3) {
			case 0:
				fmt.Println("You win, dick!")
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
