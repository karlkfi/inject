package inject

import (
	"fmt"
	"reflect"
)

type Definition interface {
	Ptr() interface{}
	Resolve(Graph) reflect.Value
	fmt.Stringer
}

type definition struct {
	ptr      interface{}
	provider Provider
	value    *reflect.Value
}

func NewDefinition(ptr interface{}, provider Provider) Definition {
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		panic("ptr is not a pointer")
	}

	targetType := reflect.ValueOf(ptr).Elem().Type()
	if !provider.ReturnType().AssignableTo(targetType) {
		panic("provider return type must be assignable to the ptr value type")
	}

	return &definition{
		ptr:      ptr,
		provider: provider,
	}
}

func (d definition) Ptr() interface{} {
	return d.ptr
}

func (d *definition) Resolve(g Graph) reflect.Value {
	if d.value != nil {
		return *d.value
	}

	value := d.provider.Provide(g)

	// cache the result
	d.value = &value

	// update the ptr value
	reflect.ValueOf(d.ptr).Elem().Set(value)

	return value
}

func (d definition) String() string {
	return fmt.Sprintf("&definition{\n%s,\n%s,\n%s\n}",
		indent(fmt.Sprintf("ptr: %s", ptrString(d.ptr)), 1),
		indent(fmt.Sprintf("provider: %s", d.provider), 1),
		indent(fmt.Sprintf("value: %s", d.value), 1),
	)
}
