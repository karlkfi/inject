package inject

import (
	"reflect"
	"fmt"
)

type graph struct {
	providers map[interface{}]Provider
	values    map[interface{}]reflect.Value
}

// NewGraph constructs a new Graph, initializing the provider and value maps.
func NewGraph() Graph {
	return &graph{
		providers: map[interface{}]Provider{},
		values:    map[interface{}]reflect.Value{},
	}
}

// Define a pointer as being resolved by a provider
func (g *graph) Define(ptr interface{}, provider Provider) {
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		panic("ptr is not a pointer")
	}

    targetType := reflect.ValueOf(ptr).Elem().Type()

	if !provider.ReturnType().AssignableTo(targetType) {
		panic("provider return type must be assignable to the ptr value type")
	}

	g.providers[ptr] = provider
}

// Resolve a type into a value by recursively resolving its dependencies and/or returning the cached result
func (g *graph) ResolveByType(ptrType reflect.Type) reflect.Value {

	var (
		found bool
		assignablePtr interface{}
	)
	for ptr := range g.providers {
		if reflect.TypeOf(ptr).Elem().AssignableTo(ptrType) {
			if found {
				panic("multiple defined pointers are assignable to the specified type")
			}
			found = true
			assignablePtr = ptr
		}
	}

	if !found {
		panic("no defined pointer is assignable to the specified type")
	}

	return g.Resolve(assignablePtr)
}

// Resolve a pointer into a value by recursively resolving its dependencies and/or returning the cached result
func (g *graph) Resolve(ptr interface{}) reflect.Value {
	value, found := g.values[ptr]
	if found {
		// value already evaluated, return the cached result
		return value
	}

	ptrType := reflect.TypeOf(ptr)
	if ptrType.Kind() != reflect.Ptr {
		panic("ptr is not a pointer")
	}

	ptrValueElem := reflect.ValueOf(ptr).Elem()
	provider, found := g.providers[ptr]
	if !found {
		// no known provider - return the current value of the pointer
		return ptrValueElem
	}

    if !provider.ReturnType().AssignableTo(ptrValueElem.Type()) {
        panic("provider return type must be assignable to the ptr value type")
    }

	value = provider.Provide(g)

	// cache the result
	g.values[ptr] = value

	// set the ptr value to the result
	reflect.ValueOf(ptr).Elem().Set(value)

	return value
}

// ResolveAll known pointers into values, caching the results
func (g *graph) ResolveAll() {
	for ptr := range g.providers {
		g.Resolve(ptr)
	}
}

// String returns a multiline string representation of the dependency graph
func (g *graph) String() string {
	return fmt.Sprintf("&graph{\n%s,\n%s\n}",
		indent(fmt.Sprintf("providers: %s", g.fmtProviders()), 1),
		indent(fmt.Sprintf("values: %s", g.fmtValues()), 1),
	)
}

func (g *graph) fmtProviders() string {
	m := make(map[string]string, len(g.providers))
	for ptr, provider := range g.providers {
		m[ptrString(ptr)] = provider.String()
	}
	return mapString(m)
}

func (g *graph) fmtValues() string {
	m := make(map[string]string, len(g.values))
	for ptr, value := range g.values {
		m[ptrString(ptr)] = value.String()
	}
	return mapString(m)
}
