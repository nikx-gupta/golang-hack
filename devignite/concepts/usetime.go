package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().UTC())
	fmt.Println(time.Parse("02 Mon 2020 22:10", "02 Tue 2021 22:21"))
}
