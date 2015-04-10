package inject

import (
	"fmt"
	"reflect"
)

type autoProvider struct {
	constructor interface{}
}

// NewAutoProvider specifies how to construct a value given its constructor function.
// Argument values are auto-resolved by type.
func NewAutoProvider(constructor interface{}) Provider {
	fnValue := reflect.ValueOf(constructor)
	if fnValue.Kind() != reflect.Func {
		panic("constructor is not a function")
	}

	fnType := reflect.TypeOf(constructor)
	if fnType.NumOut() != 1 {
		panic("constructor must have exactly 1 return value")
	}

	return autoProvider{
		constructor: constructor,
	}
}

// Provide returns the result of executing the constructor with argument values resolved by type from a dependency graph
func (p autoProvider) Provide(g Graph) reflect.Value {
	fnType := reflect.TypeOf(p.constructor)

	argCount := fnType.NumIn()
	args := make([]reflect.Value, argCount, argCount)
	for i := 0; i < argCount; i++ {
		args[i] = g.ResolveByType(fnType.In(i))
	}

	return reflect.ValueOf(p.constructor).Call(args)[0]
}

// Type returns the type of value to expect from Provide
func (p autoProvider) ReturnType() reflect.Type {
	return reflect.TypeOf(p.constructor).Out(0)
}

// String returns a multiline string representation of the autoProvider
func (p autoProvider) String() string {
	return fmt.Sprintf("&autoProvider{\n%s\n}",
		indent(fmt.Sprintf("constructor: %s", reflect.TypeOf(p.constructor)), 1),
	)
}
