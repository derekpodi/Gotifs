package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/bitfield/script"
)

func main() {

	//Read the contecnts of a file as a string
	contents, err := script.File("test.txt").String()

	//Count the lines in that file
	numLines, err := script.File("test.txt").CountLines()

	//Count the number of lines in the file that match the string "Error":
	numErrors, err := script.File("test.txt").Match("Error").CountLines()

	//pipe input into this program, and have it output only matching lines (like grep)
	script.Stdin().Match("Error").Stdout()

	//Pass results through Go function
	script.Stdin().Match("Error").FilterLine(strings.ToUpper).Stdout()

	//Pass in a list of files on the command line, and have our program read them all in sequence and output the matching lines:
	script.Args().Concat().Match("Error").Stdout()

	//First 10 matches only
	script.Args().Concat().Match("Error").First(10).Stdout()

	//Append that output to a file instead of printing it to the terminal
	script.Args().Concat().Match("Error").First(10).AppendFile("/var/log/errors.txt")

	//execute some external program instead
	script.Exec("ping 127.0.0.1").Stdout()

	//we don't know the arguments yet; we might get them from the user, for example. We'd like to be able to run the external command repeatedly, each time passing it the next line of input
	script.Args().ExecForEach("ping -c 1 {{.}}").Stdout()

	script.Echo("hello world").Filter(func(r io.Reader, w io.Writer) error {
		n, err := io.Copy(w, r)
		fmt.Fprintf(w, "\nfiltered %d bytes\n", n)
		return err
	}).Stdout()
	// Output:
	// hello world
	// filtered 11 bytes

	script.Echo("a\nb\nc").FilterScan(func(line string, w io.Writer) {
		fmt.Fprintf(w, "scanned line: %q\n", line)
	}).Stdout()
	// Output:
	// scanned line: "a"
	// scanned line: "b"
	// scanned line: "c"

	//Realistic Use case - Apache log format -- count most frequent visitors
	script.Stdin().Column(1).Freq().First(10).Stdout()
	//go run main.go <access.log
}
