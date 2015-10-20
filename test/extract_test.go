package test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/karlkfi/inject"
)


func TestExtractByType(t *testing.T) {
	RegisterTestingT(t)

	graph := inject.NewGraph()

	var (
		a1 *alpha
		b1 *beta
		g1 *gamma
	)

	graph.Define(&a1, inject.NewProvider(func() *alpha{ return &alpha{name: "a1"} } ))
	graph.Define(&b1, inject.NewProvider(func() *beta{ return &beta{name: "b1"} } ))
	graph.Define(&g1, inject.NewProvider(func() *gamma{ return &gamma{name: "g1"} } ))

	var a *alpha
	inject.ExtractByType(graph, &a)

	Expect(a).To(Equal(&alpha{name: "a1"}))
}

func TestExtractByTypeNoMatch(t *testing.T) {
	RegisterTestingT(t)

	graph := inject.NewGraph()

	var (
		b1 *beta
		g1 *gamma
	)

	graph.Define(&b1, inject.NewProvider(func() *beta{ return &beta{name: "b1"} } ))
	graph.Define(&g1, inject.NewProvider(func() *gamma{ return &gamma{name: "g1"} } ))

	var a *alpha

	defer ExpectPanic("no defined pointer matches the specified type")
	inject.ExtractByType(graph, &a)
}

func TestExtractByTypeMultiMatch(t *testing.T) {
	RegisterTestingT(t)

	graph := inject.NewGraph()

	var (
		a1 *alpha
		a2 *alpha
		b1 *beta
		g1 *gamma
	)

	graph.Define(&a1, inject.NewProvider(func() *alpha{ return &alpha{name: "a1"} } ))
	graph.Define(&a2, inject.NewProvider(func() *alpha{ return &alpha{name: "a2"} } ))
	graph.Define(&b1, inject.NewProvider(func() *beta{ return &beta{name: "b1"} } ))
	graph.Define(&g1, inject.NewProvider(func() *gamma{ return &gamma{name: "g1"} } ))

	var a *alpha

	defer ExpectPanic("more than one defined pointer matches the specified type")
	inject.ExtractByType(graph, &a)
}

func TestExtractAssignable(t *testing.T) {
	RegisterTestingT(t)

	graph := inject.NewGraph()

	var (
		a1 *alpha
		g1 *gamma
	)

	graph.Define(&a1, inject.NewProvider(func() *alpha{ return &alpha{name: "a1"} } ))
	graph.Define(&g1, inject.NewProvider(func() *gamma{ return &gamma{name: "g1"} } ))

	var o omega
	inject.ExtractAssignable(graph, &o)

	Expect(o).To(Equal(&alpha{name: "a1"}))
}

func TestExtractAssignableNoMatch(t *testing.T) {
	RegisterTestingT(t)

	graph := inject.NewGraph()

	var (
		g1 *gamma
	)

	graph.Define(&g1, inject.NewProvider(func() *gamma{ return &gamma{name: "g1"} } ))

	var o omega

	defer ExpectPanic("no defined pointer is assignable to the specified type")
	inject.ExtractAssignable(graph, &o)
}

func TestExtractAssignableMultiMatch(t *testing.T) {
	RegisterTestingT(t)

	graph := inject.NewGraph()

	var (
		a1 *alpha
		b1 *beta
		g1 *gamma
	)

	graph.Define(&a1, inject.NewProvider(func() *alpha{ return &alpha{name: "a1"} } ))
	graph.Define(&b1, inject.NewProvider(func() *beta{ return &beta{name: "b1"} } ))
	graph.Define(&g1, inject.NewProvider(func() *gamma{ return &gamma{name: "g1"} } ))

	var o omega

	defer ExpectPanic("more than one defined pointer is assignable to the specified type")
	inject.ExtractAssignable(graph, &o)
}

func ExpectPanic(content string) {
	Expect(recover()).To(ContainSubstring(content))
}
