package inject

import (
	"fmt"
)

type InterfaceC interface {
	C() string
	fmt.Stringer
}

type implC struct {}

func NewC() InterfaceC {
	return &implC{}
}

func (c *implC) C() string {
	return "C()"
}

func (c *implC) String() string {
	return "&implC{}"
}
