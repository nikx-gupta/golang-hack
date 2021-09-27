// Illustrate STD IO commands
// package should be named main so that we can run from command line
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// `os` package, as name indicates, allows access to some of the underlying os properties liek

	readStdInputAsFile()
}

func readStdInputAsFile() {

	// As each Std stream is a file, w can read it using File functions
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}

func writeOutput() {
	// Enable writing writer which can be stdOUT or any other `Writer` interface
	io.WriteString(os.Stdout, "This is written to os.Stdout stream")
}

func writeArguments() {
	// Getting Args passed to the program as String[]
	args := os.Args
	for i := 0; i < len(args); i++ {
		fmt.Printf("Argument: %d -> %s \n", i, args[i])
	}
}
