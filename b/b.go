package b

import (
	"github.com/ximyro/multi-module-repo/a"
)

type B struct {
	a.A
}

func (b *B) CallA() {}

func NewB() *B {
	return &B{}
}
