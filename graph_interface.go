package inject

import (
	"reflect"
	"fmt"
)

// Graph describes a dependency graph that resolves nodes using well defined relationships.
// These relationships are defined with node pointers and Providers.
type Graph interface {
	Define(ptr interface{}, provider Provider)
	ResolveByType(ptrType reflect.Type) reflect.Value
	Resolve(ptr interface{}) reflect.Value
	ResolveAll()
	fmt.Stringer
}
