package inject

import (
	"fmt"
	"reflect"
)

type Definition interface {
	Ptr() interface{}
	Resolve(Graph) reflect.Value
	Obscure(g Graph)
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
		panic(fmt.Sprintf("provider return type (%v) must be assignable to the ptr value type (%v)", provider.ReturnType(), targetType))
	}

	return &definition{
		ptr:      ptr,
		provider: provider,
	}
}

func (d definition) Ptr() interface{} {
	return d.ptr
}

// Resolve calls the provider, initializes the result, and populates the pointer with the result value
func (d *definition) Resolve(g Graph) reflect.Value {
	if d.value != nil {
		// already resolved
		return *d.value
	}

	value := d.provider.Provide(g)

	obj, ok := value.Interface().(Initializable)
	if ok && obj != nil {
		obj.Initialize()
	}

	// cache the result
	d.value = &value

	// update the ptr value
	reflect.ValueOf(d.ptr).Elem().Set(value)

	return value
}

// Obscure zeros out the pointer value and finalizes its previous value
func (d *definition) Obscure(g Graph) {
	if d.value == nil {
		// already obscured
		return
	}

	obj, ok := d.value.Interface().(Finalizable)

	// uncache the result
	d.value = nil

	// zero out the ptr value
	ptrValue := reflect.ValueOf(d.ptr).Elem()
	ptrValue.Set(reflect.Zero(ptrValue.Type()))

	if ok && obj != nil {
		obj.Finalize()
	}
}

func (d definition) String() string {
	return fmt.Sprintf("&definition{\n%s,\n%s,\n%s\n}",
		indent(fmt.Sprintf("ptr: %s", ptrString(d.ptr)), 1),
		indent(fmt.Sprintf("provider: %s", d.provider), 1),
		indent(fmt.Sprintf("value: %s", d.value), 1),
	)
}
