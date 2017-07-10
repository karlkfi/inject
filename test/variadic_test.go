package test

import (
	"testing"

	"github.com/karlkfi/inject"
	"github.com/mumia/inject/test/variadic"
	. "github.com/onsi/gomega"
)

func TestVariadicNone(t *testing.T) {
	RegisterTestingT(t)

	graph := inject.NewGraph()

	var container *variadic.VariadicContainer

	graph.Define(
		&container,
		inject.NewProvider(
			func() *variadic.VariadicContainer{
				return variadic.NewVariadicContainer()
			},
		),
	)
	graph.ResolveAll()

	Expect(container.GetInstalled()).To(Equal(""))
}

func TestVariadicOne(t *testing.T) {
	RegisterTestingT(t)

	graph := inject.NewGraph()

	var container *variadic.VariadicContainer

	graph.Define(
		&container,
		inject.NewProvider(
			func() *variadic.VariadicContainer{
				return variadic.NewVariadicContainer(variadic.NewV1())
			},
		),
	)
	graph.ResolveAll()

	Expect(container.GetInstalled()).To(Equal(",v1"))
}

func TestVariadicAll(t *testing.T) {
	RegisterTestingT(t)

	graph := inject.NewGraph()

	var container *variadic.VariadicContainer

	graph.Define(
		&container,
		inject.NewProvider(
			func() *variadic.VariadicContainer{
				return variadic.NewVariadicContainer(
					variadic.NewV1(),
					variadic.NewV2(),
					variadic.NewV3(),
				)
			},
		),
	)
	graph.ResolveAll()

	Expect(container.GetInstalled()).To(Equal(",v1,v2,v3"))
}
