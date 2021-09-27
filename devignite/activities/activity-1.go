// Illustrate few basic programs to understand go world
// package should be named main so that we can run from command line
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	findSum()
}

func findSum() {
	// Find sum of all the command line arguments that are valid numbers
	var stdInFile *os.File
	stdInFile = os.Stdin
	scanner := bufio.NewScanner(stdInFile)
	sumInt := 0
	for scanner.Scan() {
		input := scanner.Text()
		if input == "done" {
			break
		}

		objNum, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal("Invalid Integer Argument: \n", err)
			os.Exit(0)
		}
		sumInt += objNum
	}

	fmt.Printf("The Final sum is %d \n", sumInt)
}
