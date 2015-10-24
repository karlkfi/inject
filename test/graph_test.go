package test

import (
	"fmt"
	"testing"

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
  definitions: \[
    &definition\{
      ptr: \*test\.InterfaceC=0x.*,
      provider: &provider\{
        constructor: func\(\) test\.InterfaceC,
        argPtrs: \[\]
      \},
      value: <test\.InterfaceC Value>
    \}
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
  definitions: \[
    &definition\{
      ptr: \*\*test\.ImplD=0x.*,
      provider: &provider\{
        constructor: func\(\) \*test\.ImplD,
        argPtrs: \[\]
      \},
      value: <\*test\.ImplD Value>
    \}
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
  definitions: \[
    &definition\{
      ptr: \*test\.InterfaceA=0x.*,
      provider: &provider\{
        constructor: func\(test\.InterfaceB\) test\.InterfaceA,
        argPtrs: \[
          \*test\.InterfaceB=0x.*
        \]
      \},
      value: <test\.InterfaceA Value>
    \},
    &definition\{
      ptr: \*test\.InterfaceB=0x.*,
      provider: &provider\{
        constructor: func\(string\) test\.InterfaceB,
        argPtrs: \[
          \*string=0x.*
        \]
      \},
      value: <test\.InterfaceB Value>
    \}
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
  definitions: \[
    &definition\{
      ptr: \*test\.InterfaceA=0x.*,
      provider: &autoProvider\{
        constructor: func\(test\.InterfaceB\) test\.InterfaceA
      \},
      value: <test\.InterfaceA Value>
    },
    &definition\{
      ptr: \*test\.InterfaceB=0x.*,
      provider: &provider\{
        constructor: func\(string\) test\.InterfaceB,
        argPtrs: \[
          \*string=0x.*
        \]
      \},
      value: <test\.InterfaceB Value>
    \}
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
  definitions: \[
    &definition\{
      ptr: \*fmt\.Stringer=0x.*,
      provider: &provider\{
        constructor: func\(\) \*test\.ImplD,
        argPtrs: \[\]
      \},
      value: <\*test\.ImplD Value>
    \}
  \]
\}`
	Expect(graph.String()).To(MatchRegexp(expectedString))
}

func TestGraphSupportsPartialResolution(t *testing.T) {
	RegisterTestingT(t)

	var (
		name = "FullName"
		a    InterfaceA
		b    InterfaceB
	)

	graph := inject.NewGraph(
		inject.NewDefinition(&a, inject.NewProvider(NewA, &b)),
		inject.NewDefinition(&b, inject.NewProvider(NewB, &name)),
	)

	Expect(a).To(BeNil())
	Expect(b).To(BeNil())

	graph.Resolve(&b)

	Expect(a).To(BeNil())

	Expect(b).To(Equal(NewB(name)))
	Expect(b.String()).To(Equal("&implB{name: \"FullName\"}"))

	expectedString := `&graph\{
  definitions: \[
    &definition\{
      ptr: \*test\.InterfaceA=0x.*,
      provider: &provider\{
        constructor: func\(test\.InterfaceB\) test\.InterfaceA,
        argPtrs: \[
          \*test\.InterfaceB=0x.*
        \]
      \},
      value: <nil>
    \},
    &definition\{
      ptr: \*test\.InterfaceB=0x.*,
      provider: &provider\{
        constructor: func\(string\) test\.InterfaceB,
        argPtrs: \[
          \*string=0x.*
        \]
      \},
      value: <test\.InterfaceB Value>
    \}
  \]
\}`
	Expect(graph.String()).To(MatchRegexp(expectedString))
}

func TestGraphLifecycle(t *testing.T) {
  RegisterTestingT(t)

  var (
    i *initme
    f *finalme
    l *lifecycleme
  )

  graph := inject.NewGraph(
    inject.NewDefinition(&i, inject.NewProvider(func() *initme { return &initme{} })),
    inject.NewDefinition(&f, inject.NewProvider(func() *finalme { return &finalme{} })),
    inject.NewDefinition(&l, inject.NewProvider(func() *lifecycleme { return &lifecycleme{} })),
  )

  Expect(i).To(BeNil())
  Expect(f).To(BeNil())
  Expect(l).To(BeNil())

  graph.ResolveAll()

  // defined pointer values will be constructed and initialized
  Expect(i).To(Equal(&initme{initialized: true}))
  Expect(f).To(Equal(&finalme{finalized: false}))
  Expect(l).To(Equal(&lifecycleme{initialized: true, finalized: false}))

  ii := i
  ff := f
  ll := l

  graph.Finalize()

  // defined pointers will be zeroed
  Expect(i).To(BeNil())
  Expect(f).To(BeNil())
  Expect(l).To(BeNil())

  // values pointed at by defined pointers will be finalized
  Expect(ii).To(Equal(&initme{initialized: true}))
  Expect(ff).To(Equal(&finalme{finalized: true}))
  Expect(ll).To(Equal(&lifecycleme{initialized: true, finalized: true}))
}
