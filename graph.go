package inject

import (
	"reflect"
)

type Graph interface {
	Define(ptr interface{}, provider Provider)
	Resolve(ptr interface{}) reflect.Value
	ResolveAll()
}

type graph struct {
	providers map[interface{}]Provider
	values    map[interface{}]reflect.Value
}

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

	if provider.ConstructorType().Out(0).Kind() != reflect.ValueOf(ptr).Elem().Kind() {
		panic("constructor return value type must match ptr value type")
	}

	g.providers[ptr] = provider
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

	if provider.ConstructorType().Out(0).Kind() != ptrValueElem.Kind() {
		panic("constructor return value type must match ptr value type")
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
