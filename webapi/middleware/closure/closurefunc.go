package main

import "fmt"

func NumGenFunc() func() int {
	var i = 0
	return func() int {
		i++
		return i
	}
}

func main() {
	numgen1 := NumGenFunc()
	numgen2 := NumGenFunc()
	for i := 0; i < 5; i++ {
		fmt.Print("first function - ", numgen1(), "\t")
		fmt.Print("second function - ", numgen2(), "\n")
	}
}
