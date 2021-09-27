package main

import (
	"fmt"
)

func main() {

	//arr_1 := [5]int{1, 2, 3, 4, 5}
	slice_1 := []int{1, 2, 3, 4, 5, 6, 7}

	slice_1 = append(slice_1, 100)
	slice_1 = append(slice_1, 100)
	slice_1 = append(slice_1, 100)

	//for index, item := range slice_1 {
	//	fmt.Printf("Index - %d, Value - %d\n", index, item)
	//}

	// Get items from 1 - 2
	part_1 := slice_1[1:3]
	for index, item := range part_1 {
		fmt.Printf("Index - %d, Value - %d\n", index, item)
	}

	// Get items starting from 3 till end
	part_2 := slice_1[4:]
	for index, item := range part_2 {
		fmt.Printf("Index - %d, Value - %d\n", index, item)
	}

	// Change Slice value
	part_2[0] = 400
	fmt.Printf("Index - %d, Value - %d\n", 4, slice_1[4])
	fmt.Printf("Index - %d, Value - %d\n", 4, part_2[0])

	// Copy to larger Slice
	slice_3 := []int{1, 2, 3, 4, 5, 6}
	slice_4_cp := []int{0, 0, 0, 0, 0, 0, 0, 0}
	copy(slice_4_cp, slice_3)
	for index, item := range slice_4_cp {
		fmt.Printf("Index - %d, Value - %d\n", index, item)
	}

	fmt.Println("This is Copying to shorter slice")
	// Copy to shorter Slice
	slice_5_cp := []int{0, 0, 0}
	copy(slice_5_cp, slice_3)
	for index, item := range slice_5_cp {
		fmt.Printf("Index - %d, Value - %d\n", index, item)
	}
}
