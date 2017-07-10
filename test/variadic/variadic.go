package variadic

import "fmt"

type Variadic interface {
	Something() bool
	fmt.Stringer
}

type VariadicContainer struct {
	list []Variadic
}

func NewVariadicContainer(list ...Variadic) *VariadicContainer {
	return &VariadicContainer{list: list}
}

func (vc VariadicContainer) GetInstalled() string {
	var installed string

	for _, item := range vc.list {
		installed = fmt.Sprintf("%s,%s", installed, item.String())
	}

	return installed
}
