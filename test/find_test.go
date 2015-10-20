package test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/karlkfi/inject"
)

func TestFindByType(t *testing.T) {
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

	var aList []*alpha
	inject.FindByType(graph, &aList)

	// alphas but not betas or gammas
	Expect(aList).To(ConsistOf(&alpha{name: "a1"}, &alpha{name: "a2"}))
}

func TestFindAssignable(t *testing.T) {
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

	var oList []omega
	inject.FindAssignable(graph, &oList)

	// alphas and betas (omegas), but not gammas
	Expect(oList).To(ConsistOf(&alpha{name: "a1"}, &alpha{name: "a2"}, &beta{name: "b1"}))
}
