package inject

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGraphSupportsInterfaces(t *testing.T) {
	RegisterTestingT(t)

	graph := NewGraph()

	var (
		c InterfaceC
	)

	graph.Define(&c, NewProvider(NewC))
	graph.ResolveAll()

	Expect(c).To(Equal(NewC()))
	Expect(c.String()).To(Equal("&implC{}"))

	expectedString := `&graph\{
  providers: map\[
    \*inject\.InterfaceC=0x.*: &provider\{
      constructor: func\(\) inject\.InterfaceC,
      argPtrs: \[\]
    \}
  \],
  values: map\[
    \*inject\.InterfaceC=0x.*: <inject\.InterfaceC Value>
  \]
\}`
	Expect(graph.String()).To(MatchRegexp(expectedString))
}

func TestGraphSupportsStructPointers(t *testing.T) {
	RegisterTestingT(t)

	graph := NewGraph()

	var (
		d *ImplD
	)

	graph.Define(&d, NewProvider(NewD))
	graph.ResolveAll()

	Expect(d).To(Equal(NewD()))
	Expect(d.String()).To(Equal("&ImplD{}"))

	expectedString := `&graph\{
  providers: map\[
    \*\*inject\.ImplD=0x.*: &provider\{
      constructor: func\(\) \*inject\.ImplD,
      argPtrs: \[\]
    \}
  \],
  values: map\[
    \*\*inject\.ImplD=0x.*: <\*inject\.ImplD Value>
  \]
\}`
	Expect(graph.String()).To(MatchRegexp(expectedString))
}

func TestGraphSupportsConstructorArgs(t *testing.T) {
	RegisterTestingT(t)

	graph := NewGraph()

	var (
		name = "FullName"
		a    InterfaceA
		b    InterfaceB
	)

	graph.Define(&a, NewProvider(NewA, &b))
	graph.Define(&b, NewProvider(NewB, &name))
	graph.ResolveAll()

	Expect(a).To(Equal(NewA(NewB(name))))
	Expect(a.String()).To(Equal("&implA{b: &implB{name: \"FullName\"}}"))

	Expect(b).To(Equal(NewB(name)))
	Expect(b.String()).To(Equal("&implB{name: \"FullName\"}"))

	expectedString := `&graph\{
  providers: map\[
    \*inject\.InterfaceA=0x.*: &provider\{
      constructor: func\(inject\.InterfaceB\) inject\.InterfaceA,
      argPtrs: \[
        \*inject\.InterfaceB=0x.*
      \]
    \},
    \*inject\.InterfaceB=0x.*: &provider\{
      constructor: func\(string\) inject\.InterfaceB,
      argPtrs: \[
        \*string=0x.*
      \]
    \}
  \],
  values: map\[
    \*inject\.InterfaceA=0x.*: <inject\.InterfaceA Value>,
    \*inject\.InterfaceB=0x.*: <inject\.InterfaceB Value>
  \]
\}`
	Expect(graph.String()).To(MatchRegexp(expectedString))
}
