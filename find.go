package inject

import (
	"fmt"
	"reflect"
)

// FindByType resolves all defined pointers that match the type of the supplied slice
// and appends the resolved values to the slice.
func FindByType(g Graph, listPtr interface{}) []reflect.Value {
	ptrType := reflect.TypeOf(listPtr)
	if ptrType.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("listPtr (%v) is not a pointer", ptrType))
	}

	listType := ptrType.Elem()
	if listType.Kind() != reflect.Slice {
		panic(fmt.Sprintf("listPtr (%v) is not a pointer to a slice or array", ptrType))
	}

	listValue := reflect.ValueOf(listPtr).Elem()

	values := g.ResolveByType(listType.Elem())
	listValue = reflect.Append(listValue, values...)

	// update the listPtr value
	reflect.ValueOf(listPtr).Elem().Set(listValue)

	return values
}

// FindAssignable resolves all defined pointers that are assignable to the type of the supplied slice
// and appends the resolved values to the slice.
func FindAssignable(g Graph, listPtr interface{}) []reflect.Value {
	ptrType := reflect.TypeOf(listPtr)
	if ptrType.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("listPtr (%v) is not a pointer", ptrType))
	}

	listType := ptrType.Elem()
	if listType.Kind() != reflect.Slice {
		panic(fmt.Sprintf("listPtr (%v) is not a pointer to a slice or array", ptrType))
	}

	listValue := reflect.ValueOf(listPtr).Elem()

	values := g.ResolveByAssignableType(listType.Elem())
	listValue = reflect.Append(listValue, values...)

	// update the listPtr value
	reflect.ValueOf(listPtr).Elem().Set(listValue)

	return values
}
