package inject

import (
	"fmt"
)

type InterfaceB interface {
	B() string
	fmt.Stringer
}

type implB struct {
	name string
}

func NewB(name string) InterfaceB {
	return &implB{
		name: name,
	}
}

func (b *implB) B() string {
	return fmt.Sprintf("B() -> %s", b.name)
}

func (b *implB) String() string {
	return fmt.Sprintf("&implB{name: %#v}", b.name)
}
