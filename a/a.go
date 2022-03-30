package a

import "fmt"

type A struct{}

func (a *A) CallA() {
	fmt.Println("Calling A")
}

func NewA() *A {
	return &A{}
}
