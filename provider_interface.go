package inject

import (
	"reflect"
)

// Provider describes how to retrieve (or construct) a generic value, given a dependency graph.
type Provider interface {
	Kind() reflect.Kind
	Provide(Graph) reflect.Value
}
