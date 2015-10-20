package inject

import (
	"fmt"
	"reflect"
)

// ExtractByType resolves a pointer into a value by finding exactly one defined pointer with the specified type
func ExtractByType(g Graph, ptr interface{}) reflect.Value {
	ptrType := reflect.TypeOf(ptr)
	if ptrType.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("ptr (%v) is not a pointer", ptrType))
	}

	targetType := reflect.ValueOf(ptr).Elem().Type()
	values := g.ResolveByType(targetType)

	if len(values) > 1 {
		panic(fmt.Sprintf("more than one defined pointer matches the specified type (%v)", ptr))
	} else if len(values) == 0 {
		panic(fmt.Sprintf("no defined pointer matches the specified type (%v)", ptr))
	}
	value := values[0]

	// update the ptr value
	reflect.ValueOf(ptr).Elem().Set(value)

	return value
}

// ExtractAssignable resolves a pointer into a value by finding exactly one defined pointer with an assignable type
func ExtractAssignable(g Graph, ptr interface{}) reflect.Value {
	ptrType := reflect.TypeOf(ptr)
	if ptrType.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("ptr (%v) is not a pointer", ptrType))
	}

	targetType := reflect.ValueOf(ptr).Elem().Type()
	values := g.ResolveByAssignableType(targetType)

	if len(values) > 1 {
		panic(fmt.Sprintf("more than one defined pointer is assignable to the specified type (%v)", ptr))
	} else if len(values) == 0 {
		panic(fmt.Sprintf("no defined pointer is assignable to the specified type (%v)", ptr))
	}
	value := values[0]

	// update the ptr value
	reflect.ValueOf(ptr).Elem().Set(value)

	return value
}
