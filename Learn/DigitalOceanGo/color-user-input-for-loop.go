//https://www.digitalocean.com/community/tutorials/how-to-write-comments-in-go

package main

import (
	"fmt"
	"strings"
)

const favColor string = "BLUE"

func main() {
	var guess string
	// Create an input loop
	for {
		// Ask the user to guess my favorite color
		fmt.Println("Guess my favorite color:")
		// Try to read a line of input from the user. Print out the error 0
		if _, err := fmt.Scanln(&guess); err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		// Did they guess the correct color?
		if favColor == strings.ToTitle(guess) {
			// They guessed it!
			fmt.Printf("%q is my favorite color!\n", favColor)
			return
		}
		// Wrong! Have them guess again.
		fmt.Printf("Sorry, %q is not my favorite color. Guess again.\n", guess)
	}
}
