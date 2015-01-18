package inject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphSupportsInterfaces(t *testing.T) {
	graph := NewGraph()

	var (
		c InterfaceC
	)

	graph.Define(&c, NewProvider(NewC))
	graph.ResolveAll()

	assert.Equal(t, NewC(), c)

	assert.Equal(t, "&implC{}", c.String())
}

func TestGraphSupportsStructPointers(t *testing.T) {
	graph := NewGraph()

	var (
		d *ImplD
	)

	graph.Define(&d, NewProvider(NewD))
	graph.ResolveAll()

	assert.Equal(t, NewD(), d)

	assert.Equal(t, "&ImplD{}", d.String())
}

func TestGraphSupportsConstructorArgs(t *testing.T) {
	graph := NewGraph()

	var (
		name = "FullName"
		a    InterfaceA
		b    InterfaceB
	)

	graph.Define(&a, NewProvider(NewA, &b))
	graph.Define(&b, NewProvider(NewB, &name))
	graph.ResolveAll()

	assert.Equal(t, NewA(NewB(name)), a)
	assert.Equal(t, NewB(name), b)

	assert.Equal(t, "&implA{b: &implB{name: \"FullName\"}}", a.String())
	assert.Equal(t, "&implB{name: \"FullName\"}", b.String())
}
