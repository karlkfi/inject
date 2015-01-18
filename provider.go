package inject

import (
	"reflect"
)

type Provider interface {
	ConstructorType() reflect.Type
	Provide(Graph) reflect.Value
}

type provider struct {
	constructor interface{}
	argPtrs     []interface{}
}

// NewProvider specifies how to construct a value given its constructor function and argument pointers
func NewProvider(constructor interface{}, argPtrs ...interface{}) Provider {
	fnValue := reflect.ValueOf(constructor)
	if fnValue.Kind() != reflect.Func {
		panic("constructor is not a function")
	}

	fnType := reflect.TypeOf(constructor)
	if fnType.NumOut() != 1 {
		panic("constructor must have exactly 1 return value")
	}

	argCount := fnType.NumIn()
	if argCount != len(argPtrs) {
		panic("argPtrs must match constructor arguments")
	}

	for i := 0; i < argCount; i++ {
		//todo: validate that argPtrs[i] is a pointer
		//todo: validate that reflect.ValueOf(argPtrs[i]).Kind() matches fnType.In(i).Kind()
	}

	return &provider{
		constructor: constructor,
		argPtrs:     argPtrs,
	}
}

// Provide returns the result of executing the constructor with argument values resolved from a dependency graph
func (p *provider) Provide(g Graph) reflect.Value {
	fnType := reflect.TypeOf(p.constructor)

	argCount := fnType.NumIn()
	args := make([]reflect.Value, argCount, argCount)
	for i := 0; i < argCount; i++ {
		args[i] = g.Resolve(p.argPtrs[i])
	}
	return reflect.ValueOf(p.constructor).Call(args)[0]
}

func (p *provider) ConstructorType() reflect.Type {
	return reflect.TypeOf(p.constructor)
}
