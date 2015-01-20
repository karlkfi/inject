package inject

import (
	"reflect"
	"fmt"
)

// Provider describes how to retrieve (or construct) a generic value, given a dependency graph.
type Provider interface {
	Type() reflect.Type
	Provide(Graph) reflect.Value
	fmt.Stringer
}
