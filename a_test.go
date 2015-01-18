package inject

import (
	"fmt"
)

type InterfaceA interface {
	A() string
	fmt.Stringer
}

type implA struct {
	b InterfaceB
}

func NewA(b InterfaceB) InterfaceA {
	return &implA{
		b: b,
	}
}

func (a *implA) A() string {
	return fmt.Sprintf("A() -> %s", a.b.B())
}

func (a *implA) String() string {
	return fmt.Sprintf("&implA{b: %s}", a.b)
}
