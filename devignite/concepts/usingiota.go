package main

import (
	"fmt"
)

func main() {
	//fmt.Println(Monday)
	//fmt.Println(Tuesday)
	//fmt.Println(Wednesday)
	//fmt.Println(Thursday)
	//fmt.Println(Friday)

	//fmt.Println(Power2_1)
	//fmt.Println(Power2_2)
	//fmt.Println(Power2_3)

	fmt.Println(Power4_1)
	fmt.Println(Power4_2)
	fmt.Println(Power4_3)
	fmt.Println(Power4_4)

	//  100
	//10000
}

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
)

const (
	Power4_1 = 1 << (iota * 2)
	Power4_2
	Power4_3
	Power4_4
)

const (
	Power2_1 = 1 << iota
	Power2_2
	Power2_3
)

