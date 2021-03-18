// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

<<<<<<< HEAD
// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------
=======
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
>>>>>>> 2da8730746d42d21b8836a93f68cfd15cfec6d11

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
<<<<<<< HEAD
		n := rand.Intn(guess + 1)
		fmt.Println("Turn:", turn)
		if n == guess {
			if turn == 0 {
				fmt.Println("You won on your first try!")
			}
			fmt.Println("🎉  YOU WIN!")
			// if n is even, print w[0], else print w[1]
			//fmt.Println("n:", n)
			if n%2 == 0 {
				fmt.Println(w[0])
			} else {
				fmt.Println(w[1])
			}

			return
		}
	}

	fmt.Println("☠️  YOU LOST... Try again?")
=======
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
>>>>>>> 2da8730746d42d21b8836a93f68cfd15cfec6d11
}
