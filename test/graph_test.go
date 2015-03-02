package test

import (
	"testing"
    "fmt"

    . "github.com/onsi/gomega"

    "github.com/karlkfi/inject"
)

func TestGraphSupportsInterfaces(t *testing.T) {
	RegisterTestingT(t)

	graph := inject.NewGraph()

	var (
		c InterfaceC
	)

	graph.Define(&c, inject.NewProvider(NewC))
	graph.ResolveAll()

	Expect(c).To(Equal(NewC()))
	Expect(c.String()).To(Equal("&implC{}"))

	expectedString := `&graph\{
  providers: map\[
    \*test\.InterfaceC=0x.*: &provider\{
      constructor: func\(\) test\.InterfaceC,
      argPtrs: \[\]
    \}
  \],
  values: map\[
    \*test\.InterfaceC=0x.*: <test\.InterfaceC Value>
  \]
\}`
	Expect(graph.String()).To(MatchRegexp(expectedString))
}

func TestGraphSupportsStructPointers(t *testing.T) {
	RegisterTestingT(t)

	graph := inject.NewGraph()

	var (
		d *ImplD
	)

	graph.Define(&d, inject.NewProvider(NewD))
	graph.ResolveAll()

	Expect(d).To(Equal(NewD()))
	Expect(d.String()).To(Equal("&ImplD{}"))

	expectedString := `&graph\{
  providers: map\[
    \*\*test\.ImplD=0x.*: &provider\{
      constructor: func\(\) \*test\.ImplD,
      argPtrs: \[\]
    \}
  \],
  values: map\[
    \*\*test\.ImplD=0x.*: <\*test\.ImplD Value>
  \]
\}`
	Expect(graph.String()).To(MatchRegexp(expectedString))
}

func TestGraphSupportsProviderConstructorArgs(t *testing.T) {
	RegisterTestingT(t)

	graph := inject.NewGraph()

	var (
		name = "FullName"
		a    InterfaceA
		b    InterfaceB
	)

	graph.Define(&a, inject.NewProvider(NewA, &b))
	graph.Define(&b, inject.NewProvider(NewB, &name))
	graph.ResolveAll()

	Expect(a).To(Equal(NewA(NewB(name))))
	Expect(a.String()).To(Equal("&implA{b: &implB{name: \"FullName\"}}"))

	Expect(b).To(Equal(NewB(name)))
	Expect(b.String()).To(Equal("&implB{name: \"FullName\"}"))

	expectedString := `&graph\{
  providers: map\[
    \*test\.InterfaceA=0x.*: &provider\{
      constructor: func\(test\.InterfaceB\) test\.InterfaceA,
      argPtrs: \[
        \*test\.InterfaceB=0x.*
      \]
    \},
    \*test\.InterfaceB=0x.*: &provider\{
      constructor: func\(string\) test\.InterfaceB,
      argPtrs: \[
        \*string=0x.*
      \]
    \}
  \],
  values: map\[
    \*test\.InterfaceA=0x.*: <test\.InterfaceA Value>,
    \*test\.InterfaceB=0x.*: <test\.InterfaceB Value>
  \]
\}`
	Expect(graph.String()).To(MatchRegexp(expectedString))
}

func TestGraphSupportsAutoProvider(t *testing.T) {
	RegisterTestingT(t)

	graph := inject.NewGraph()

	var (
		name = "FullName"
		a    InterfaceA
		b    InterfaceB
	)

	graph.Define(&a, inject.NewAutoProvider(NewA))
	graph.Define(&b, inject.NewProvider(NewB, &name))
	graph.ResolveAll()

	Expect(a).To(Equal(NewA(NewB(name))))
	Expect(a.String()).To(Equal("&implA{b: &implB{name: \"FullName\"}}"))

	Expect(b).To(Equal(NewB(name)))
	Expect(b.String()).To(Equal("&implB{name: \"FullName\"}"))

	expectedString := `&graph\{
  providers: map\[
    \*test\.InterfaceA=0x.*: &autoProvider\{
      constructor: func\(test\.InterfaceB\) test\.InterfaceA
    \},
    \*test\.InterfaceB=0x.*: &provider\{
      constructor: func\(string\) test\.InterfaceB,
      argPtrs: \[
        \*string=0x.*
      \]
    \}
  \],
  values: map\[
    \*test\.InterfaceA=0x.*: <test\.InterfaceA Value>,
    \*test\.InterfaceB=0x.*: <test\.InterfaceB Value>
  \]
\}`
	Expect(graph.String()).To(MatchRegexp(expectedString))
}

func TestGraphSupportsDownCasting(t *testing.T) {
    RegisterTestingT(t)

    graph := inject.NewGraph()

    var (
        d fmt.Stringer
    )

    graph.Define(&d, inject.NewProvider(NewD))
    graph.ResolveAll()

    Expect(d).To(Equal(NewD()))
    Expect(d.String()).To(Equal("&ImplD{}"))

    expectedString := `&graph\{
  providers: map\[
    \*fmt\.Stringer=0x.*: &provider\{
      constructor: func\(\) \*test\.ImplD,
      argPtrs: \[\]
    \}
  \],
  values: map\[
    \*fmt\.Stringer=0x.*: <\*test\.ImplD Value>
  \]
\}`
    Expect(graph.String()).To(MatchRegexp(expectedString))
}
