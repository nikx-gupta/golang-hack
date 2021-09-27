package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	//time.Sleep(time.Second * 60)
	// fmt.Println(time.Now())

	fmt.Println(time.Now().Format("01 Jan 2020"))

	loc, _ := time.LoadLocation("Asia/Kolkata")
	fmt.Println(time.Now().In(loc))

	fmt.Println(time.Parse("01 Feb 2021", "20 Mar 1970"))
}
