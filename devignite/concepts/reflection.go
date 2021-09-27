package main

import (
	"fmt"
	"reflect"
)

func main() {

	x := "test"
	y := "test2"

	// Type of instant
	reflType := reflect.TypeOf(x)
	// Value of instant
	reflValue := reflect.ValueOf(x)
	reflAddValue := reflect.ValueOf(&x).Elem()
	fmt.Println(reflType)
	fmt.Println(reflValue)
	fmt.Println(reflAddValue)

	// Compare Two Types
	fmt.Println(reflect.TypeOf(x) == reflect.TypeOf(y))

}
