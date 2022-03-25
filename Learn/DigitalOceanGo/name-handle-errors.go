//https://www.digitalocean.com/community/tutorials/handling-errors-in-go

package main

import (
	"errors"
	"fmt"
	"strings"
)

func capitalize(name string) (string, int, error) {
	handle := func(err error) (string, int, error) { //Anonymous function reduce boilerplate errors
		return "", 0, err
	}

	if name == "" {
		return handle(errors.New("no name provided"))
	}

	return strings.ToTitle(name), len(name), nil
}

func main() {
	name, size, err := capitalize("derek")
	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	fmt.Printf("Capitalized name: %s, length: %d", name, size)
}
