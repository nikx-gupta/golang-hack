package main

import (
	"devignite/datastructures/hashes"
	"devignite/datastructures/trees"
	"fmt"
)

func main() {

	HashOperations()
	// TreeOperations()
	//lists.SimpleIntList()
	//lists.SimpleDataList()
	//lists.ReverseList()
}

func HashOperations(){
	hashes.DoHash()
}

func TreeOperations() {
	var tree = &trees.BinarySearchTree{}
	tree.InsertElement(10, 10)
	tree.InsertElement(8, 8)
	tree.InsertElement(7, 3)
	tree.InsertElement(15, 15)
	tree.InsertElement(13, 13)

	tree.InOrderTraverse(func(i int, node string) {
		fmt.Printf("%d - %s \n", i, node)
	})

	fmt.Printf("%d", *tree.MinNode())
}
