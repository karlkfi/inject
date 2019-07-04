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
	if !fnValue.Type().IsVariadic() && argCount != len(argPtrs) {
		panic(fmt.Sprintf("argPtrs (%d) must match constructor arguments (%d)", len(argPtrs), argCount))
	}

	var kind reflect.Kind
	for i, argPtr := range argPtrs {
		isVariadic := fnValue.Type().IsVariadic() && (fnType.NumIn() == 1 || i >= fnType.NumIn())

		if i < fnType.NumIn() {
			kind = fnType.In(i).Kind()
		} else {
			kind = fnType.In(fnType.NumIn() - 1).Kind()
		}

		if reflect.TypeOf(argPtr).Kind() != reflect.Ptr {
			panic(fmt.Sprintf("argPtrs must all be pointers, found %v", reflect.TypeOf(argPtr)))
		}

		if !isVariadic && reflect.ValueOf(argPtr).Elem().Kind() != kind {
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
	if fnType.IsVariadic() {
		argCount = len(p.argPtrs)
	}

	args := make([]reflect.Value, argCount, argCount)
	var inType reflect.Type
	for i := 0; i < argCount; i++ {
		arg := g.Resolve(p.argPtrs[i])
		argType := arg.Type()

		if fnType.IsVariadic() && i >= fnType.NumIn()-1 {
			inType = fnType.In(fnType.NumIn() - 1).Elem()
		} else {
			inType = fnType.In(i)
		}

		if !argType.AssignableTo(inType) {
			if !argType.ConvertibleTo(inType) {
				panic(fmt.Sprintf(
					"arg %d of type %q cannot be assigned or converted to type %q for provider constructor (%s)",
					i, argType, inType, p.constructor,
				))
			}
			arg = arg.Convert(inType)
		}
		args[i] = arg
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
