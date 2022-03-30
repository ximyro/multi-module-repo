package a

import "fmt"

type A struct{}

func (a *A) CallA() {
	fmt.Println("Calling new version A")
}

func NewA() *A {
	return &A{}
}
