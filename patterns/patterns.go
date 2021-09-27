package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//var obj IProgress = DerivedObject{}

	json.Marshal()
}

type IProgress interface {
	Report()
	SecondReport()
}

type IList interface {

}

type DerivedObject struct {
	list IList
}

func (t *DerivedObject) Report() {
	fmt.Println("This is Derived Report")
}

