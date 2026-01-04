package main

import "fmt"

type firstStruct struct {
}

func (*firstStruct) method() {
	fmt.Println("Выполняю что-то")
}

type Adapter interface {
	Operation()
}

type secondStruct struct {
	*firstStruct
}

func (secondStruct *secondStruct) Operation() {
	secondStruct.method()
}

func NewStruct(firstStruct *firstStruct) Adapter {
	return &secondStruct{firstStruct}
}

func main() {
	fmt.Println("Начало работы")
	interf := NewStruct(&firstStruct{})
	interf.Operation()
}
