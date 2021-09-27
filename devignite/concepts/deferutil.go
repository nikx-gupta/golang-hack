package main

import (
	"fmt"
)

func deferAnonmyousWithArgs() {
	for i := 0; i < 10; i++ {
		defer func(n int) {
			fmt.Print(n, "\n")
		}(i)
	}
}

func deferAnonmyousNoArgs() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Print(i, "\n")
		}()
	}
}

func deferPrint() {
	for i := 0; i < 10; i++ {
		defer fmt.Print(i, "\n")
	}

	fmt.Print("\n")
}

func demoPanic() {
	fmt.Println("-- Executing demoPanic --")
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Recovered when panic raised inside demoPanic")
		}
	}()

	fmt.Println("calling panic create function")
	createPanic()
	fmt.Println("Execution end demoPanic")
}

func createPanic(){
	fmt.Println("Inside createPanic")
	panic("Panic created")
	fmt.Println("End of createPanic")
}

func main() {
	// deferPrint()
	//deferAnonmyousNoArgs()
	//deferAnonmyousWithArgs()
	demoPanic()
}
