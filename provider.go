package inject

import (
	"fmt"
	"reflect"
)

type provider struct {
	constructor interface{}
	argPtrs     []interface{}
}

// NewProvider specifies how to construct a value given its constructor function and argument pointers
func NewProvider(constructor interface{}, argPtrs ...interface{}) Provider {
	fnValue := reflect.ValueOf(constructor)
	if fnValue.Kind() != reflect.Func {
		panic(fmt.Sprintf("constructor (%v) is not a function, found %v", fnValue, fnValue.Kind()))
	}

	fnType := reflect.TypeOf(constructor)
	if fnType.NumOut() != 1 {
		panic(fmt.Sprintf("constructor must have exactly 1 return value, found %v", fnType.NumOut()))
	}

	argCount := fnType.NumIn()
	if argCount != len(argPtrs) {
		panic(fmt.Sprintf("argPtrs (%d) must match constructor arguments (%d)", len(argPtrs), argCount))
	}

	for i, argPtr := range argPtrs {
		if reflect.TypeOf(argPtr).Kind() != reflect.Ptr {
			panic(fmt.Sprintf("argPtrs must all be pointers, found %v", reflect.TypeOf(argPtr)))
		}
		if reflect.ValueOf(argPtr).Elem().Kind() != fnType.In(i).Kind() {
			panic("argPtrs must match constructor argument types")
		}
	}

	return provider{
		constructor: constructor,
		argPtrs:     argPtrs,
	}
}

// Provide returns the result of executing the constructor with argument values resolved from a dependency graph
func (p provider) Provide(g Graph) reflect.Value {
	fnType := reflect.TypeOf(p.constructor)

	argCount := fnType.NumIn()
	args := make([]reflect.Value, argCount, argCount)
	for i := 0; i < argCount; i++ {
		args[i] = g.Resolve(p.argPtrs[i])
	}

	return reflect.ValueOf(p.constructor).Call(args)[0]
}

// Type returns the type of value to expect from Provide
func (p provider) ReturnType() reflect.Type {
	return reflect.TypeOf(p.constructor).Out(0)
}

// String returns a multiline string representation of the provider
func (p provider) String() string {
	return fmt.Sprintf("&provider{\n%s,\n%s\n}",
		indent(fmt.Sprintf("constructor: %s", reflect.TypeOf(p.constructor)), 1),
		indent(fmt.Sprintf("argPtrs: %s", p.fmtArgPtrs()), 1),
	)
}

func (p provider) fmtArgPtrs() string {
	b := make([]string, len(p.argPtrs), len(p.argPtrs))
	for i, argPtr := range p.argPtrs {
		b[i] = ptrString(argPtr)
	}
	return arrayString(b)
}
