package adapter

import (
	"fmt"
)

type iProcess interface {
	process()
}

type adaptee struct {
	adapterType int
}

type adapter struct {
	adaptee adaptee
}

func (adaptee adaptee) convert() {
	fmt.Println("Adaptee convert method")
}

func (adapter adapter) process() {
	fmt.Println("Adapter process")
	adapter.adaptee.convert()
}

func adapterDesign() {
	var processor iProcess = adapter{}
	processor.process()
}
