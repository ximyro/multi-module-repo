package a

type A struct{}

func (a *A) CallA() {}

func NewA() *A {
	return &A{}
}
