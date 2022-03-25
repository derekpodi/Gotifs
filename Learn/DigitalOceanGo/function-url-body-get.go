//https://www.digitalocean.com/community/tutorials/how-to-write-comments-in-go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// MustGet will retrieve a url and return the body of the page.
// If Get encounters any errors, it will panic.
func MustGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	// don't forget to close the body
	defer resp.Body.Close()
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		panic(err)
	}
	return string(body)
}

func main() {
	var url string

	// Ask the user to input url
	fmt.Println("Input your url:")
	// Try to read a line of input from the user. Print out the error 0
	if _, err := fmt.Scanln(&url); err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	body := MustGet(url)
	fmt.Printf(string(body))
}
