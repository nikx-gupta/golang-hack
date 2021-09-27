package lists

import (
	"container/list"
	"devignite/datastructures/models"
	"fmt"
)

func main() {
	//	SimpleDataList()
	ReverseList()
}


func ReverseList() {
	var intList list.List
	intList.PushFront(1)
	intList.PushFront(2)
	intList.PushFront(4)
	intList.PushFront(5)

	for elf, elb := intList.Front(), intList.Back(); elf != nil; elf, elb = elf.Next(), elb.Prev() {
		elf.Value, elb.Value = elb.Value, elf.Value
	}

	for elf := intList.Front(); elf != nil; elf = elf.Next() {
		fmt.Println(elf.Value)
	}
}

func SimpleDataList() {
	var dataList list.List
	dataList.PushFront(models.MongoRecord{
		Id: 1,
	})
	dataList.PushFront(models.MongoRecord{
		Id: 2,
	})

	for el := dataList.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value.(models.MongoRecord))
	}
}

func SimpleIntList() {
	var intList list.List

	// Initializes the List
	intList.Init()

	intList.PushFront(10)
	intList.PushFront(20)
	intList.PushFront(30)

	fmt.Println("**** Print From Front ****")

	// Traverse Front
	for el := intList.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value)
	}

	fmt.Println("**** Print From Back ****")
	// Traverse Back
	for el := intList.Back(); el != nil; el = el.Prev() {
		fmt.Println(el.Value)
	}
}
