package test

import (
	"testing"

	"github.com/karlkfi/inject"
	"github.com/karlkfi/inject/test/variadic"
	. "github.com/onsi/gomega"
)

func TestVariadicNone(t *testing.T) {
	RegisterTestingT(t)

	graph := inject.NewGraph()

	var container *variadic.VariadicContainer

	graph.Define(
		&container,
		inject.NewProvider(
			variadic.NewVariadicContainer,
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
			variadic.NewVariadicContainer,
			variadic.NewV1(),
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
			variadic.NewVariadicContainer,
			variadic.NewV1(),
			variadic.NewV2(),
			variadic.NewV3(),
		),
	)
	graph.ResolveAll()

	Expect(container.GetInstalled()).To(Equal(",v1,v2,v3"))
}

func TestNotVariadicNone(t *testing.T) {
	RegisterTestingT(t)

	graph := inject.NewGraph()

	var container *variadic.NotVariadicContainer

	graph.Define(
		&container,
		inject.NewProvider(
			variadic.NewNotVariadicContainer,
			&[]variadic.Item{},
		),
	)
	graph.ResolveAll()

	Expect(container.GetInstalled()).To(Equal(""))
}

func TestNotVariadicOne(t *testing.T) {
	RegisterTestingT(t)

	graph := inject.NewGraph()

	var container *variadic.NotVariadicContainer

	graph.Define(
		&container,
		inject.NewProvider(
			variadic.NewNotVariadicContainer,
			&[]variadic.Item{variadic.NewV1()},
		),
	)
	graph.ResolveAll()

	Expect(container.GetInstalled()).To(Equal(",v1"))
}

func TestNotVariadicAll(t *testing.T) {
	RegisterTestingT(t)

	graph := inject.NewGraph()

	var container *variadic.NotVariadicContainer

	graph.Define(
		&container,
		inject.NewProvider(
			variadic.NewNotVariadicContainer,
			&[]variadic.Item{
				variadic.NewV1(),
				variadic.NewV2(),
				variadic.NewV3(),
			},
		),
	)
	graph.ResolveAll()

	Expect(container.GetInstalled()).To(Equal(",v1,v2,v3"))
}
