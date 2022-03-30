package b

import (
	"fmt"

	"github.com/ximyro/multi-module-repo/a"
)

type B struct {
	a *a.A
}

func (b *B) CallA() {
	b.a.CallA()
	fmt.Println("Calling B")
}

func NewB(a *a.A) *B {
	return &B{
		a: a,
	}
}
