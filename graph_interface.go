package inject

import (
	"reflect"
)

// Graph describes a dependency graph that resolves nodes using well defined relationships.
// These relationships are defined with node pointers and Providers.
type Graph interface {
	Define(ptr interface{}, provider Provider)
	Resolve(ptr interface{}) reflect.Value
	ResolveAll()
}
