package variadic

import "fmt"

type Item interface {
	Something() bool
	fmt.Stringer
}

type VariadicContainer struct {
	list []Item
}

func NewVariadicContainer(list ...Item) *VariadicContainer {
	return &VariadicContainer{list: list}
}

func (vc VariadicContainer) GetInstalled() string {
	var installed string

	for _, item := range vc.list {
		installed = fmt.Sprintf("%s,%s", installed, item.String())
	}

	return installed
}

type NotVariadicContainer struct {
	list []Item
}

func NewNotVariadicContainer(list []Item) *NotVariadicContainer {
	return &NotVariadicContainer{list: list}
}

func (vc NotVariadicContainer) GetInstalled() string {
	var installed string

	for _, item := range vc.list {
		installed = fmt.Sprintf("%s,%s", installed, item.String())
	}

	return installed
}
