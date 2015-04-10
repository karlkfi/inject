package inject

import (
	"fmt"
	"reflect"
	"sort"
)

type graph struct {
	definitions map[interface{}]Definition
}

// NewGraph constructs a new Graph, initializing the provider and value maps.
func NewGraph() Graph {
	return &graph{
		definitions: map[interface{}]Definition{},
	}
}

// Define a pointer as being resolved by a provider
func (g *graph) Define(ptr interface{}, provider Provider) Definition {
	def := NewDefinition(ptr, provider, g)
	g.definitions[ptr] = def
	return def
}

// Resolve a type into a value by recursively resolving its dependencies and/or returning the cached result
func (g *graph) ResolveByType(ptrType reflect.Type) reflect.Value {
	var found Definition
	for ptr, def := range g.definitions {
		if reflect.TypeOf(ptr).Elem().AssignableTo(ptrType) {
			if found != nil {
				panic("multiple defined pointers are assignable to the specified type")
			}
			found = def
		}
	}

	if found == nil {
		panic("no defined pointer is assignable to the specified type")
	}

	return found.Resolve()
}

// Resolve a pointer into a value by recursively resolving its dependencies and/or returning the cached result
func (g *graph) Resolve(ptr interface{}) reflect.Value {
	ptrType := reflect.TypeOf(ptr)
	if ptrType.Kind() != reflect.Ptr {
		panic("ptr is not a pointer")
	}

	ptrValueElem := reflect.ValueOf(ptr).Elem()
	def, found := g.definitions[ptr]
	if !found {
		// no known definition - return the current value of the pointer
		return ptrValueElem
	}

	return def.Resolve()
}

// ResolveAll known pointers into values, caching the results
func (g *graph) ResolveAll() {
	for _, def := range g.definitions {
		def.Resolve()
	}
}

// String returns a multiline string representation of the dependency graph
func (g graph) String() string {
	return fmt.Sprintf("&graph{\n%s\n}",
		indent(fmt.Sprintf("definitions: %s", g.fmtDefinitions()), 1),
	)
}

func (g graph) fmtDefinitions() string {
	a := make([]string, 0, len(g.definitions))
	for _, def := range g.definitions {
		a = append(a, def.String())
	}
	sort.Strings(a)
	return arrayString(a)
}
