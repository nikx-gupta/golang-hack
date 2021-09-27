package main

import "fmt"

func main() {

	//arr_1 := [5]int{1, 2, 3, 4, 5}
	//slice_1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice_map := map[string]int{
		"key_1": 100,
		"key_2": 200,
	}

	// Iteration
	for key, value := range slice_map {
		fmt.Printf("Map Key: %s, Value: %d \n", key, value)
	}

	// delete item from map
	delete(slice_map, "key_1")

}
