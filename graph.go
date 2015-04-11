package inject

import (
	"fmt"
	"reflect"
	"sort"
)

// Graph describes a dependency graph that resolves nodes using well defined relationships.
// These relationships are defined with node pointers and Providers.
type Graph interface {
	Add(Definition)
	Define(ptr interface{}, provider Provider) Definition
	ResolveByType(ptrType reflect.Type) reflect.Value
	Resolve(ptr interface{}) reflect.Value
	ResolveAll()
	fmt.Stringer
}

type graph struct {
	definitions map[interface{}]Definition
}

// NewGraph constructs a new Graph, initializing the provider and value maps.
func NewGraph(defs ...Definition) Graph {
	defMap := make(map[interface{}]Definition, len(defs))
	for _, def := range defs {
		defMap[def.Ptr()] = def
	}
	return &graph{
		definitions: defMap,
	}
}

func (g *graph) Add(def Definition) {
	g.definitions[def.Ptr()] = def
}

// Define a pointer as being resolved by a provider
func (g *graph) Define(ptr interface{}, provider Provider) Definition {
	def := NewDefinition(ptr, provider)
	g.Add(def)
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

	return found.Resolve(g)
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

	return def.Resolve(g)
}

// ResolveAll known pointers into values, caching the results
func (g *graph) ResolveAll() {
	for _, def := range g.definitions {
		def.Resolve(g)
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
