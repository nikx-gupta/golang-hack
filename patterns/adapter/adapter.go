package adapter

import "fmt"

type IProcess interface {
	Process()
}

type Adapter struct {
	adaptee Adaptee
}

type Adaptee struct {
	adapterType int
}

func (adaptee Adaptee) convert() {
	fmt.Println(adaptee.adapterType)
}

func (adapter Adapter) Process() {
	fmt.Println("Adapter process")
	adapter.adaptee.convert()
}
