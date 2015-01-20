# Inject
Dependency injection library for Go (golang)

[![GoDoc](https://godoc.org/github.com/karlkfi/inject?status.svg)](https://godoc.org/github.com/karlkfi/inject)

[![Build Status](https://travis-ci.org/karlkfi/inject.svg?branch=master)](https://travis-ci.org/karlkfi/inject)

# Why use Inject?

Like most other golang injection frameworks, Inject allows auto-resolving dependencies by type.

Unlike most other golang injection frameworks, Inject uses constructor functions and argument pointers to allow
resolving any object, without it having to be designed for injection.

This has several advantages:
- Constructor arguments can be resolved automatically by type OR manually by unresolved pointer.
- Public access to the struct is not required, if they implement a public interface.
- No coordination is required between implementations of a common interface.
- Pointers can be used to distinguish between different implementations of the same interface.
- Pointers can even be used to distinguish between different instances of the same struct.

# Usage

```
package yours

import (
  "path/to/your/pkgA"
  "path/to/your/pkgB"
  "github.com/karlkfi/inject"
)

func main() {
	graph := inject.NewGraph()

	var (
		primitive = "some string"
		a    pkgA.InterfaceA
		b    *pkgB.StructB
	)

	// define how to construct a (dependent on b)
	graph.Define(&a, inject.NewProvider(pkgA.NewA, &b))

	// define how to construct b (not dependent on the graph)
	graph.Define(&b, inject.NewProvider(pkgB.NewB, &primitive))

	// resolve a and all its (transitive) dependencies
	graph.Resolve(&a)

    // a is now usable
	a.DoStuff()
}

```

The simple example above makes use of the normal providers, which require the user to specify all constructor arguments.

The even simpler example below demonstrates auto-providers, which attempt to resolve constructor arguments by only their
type. This method is simpler to use, but has the drawback that all auto-resolved types must have one and only one
pointer in the graph that is assignable to that type. Auto-provider usage can be mixed with normal provider usage for
increased flexibility.

```
...

func main() {
	graph := inject.NewGraph()

	var (
		primitive = "some string"
		a    pkgA.InterfaceA
		b    *pkgB.StructB
	)

	// define how to construct a (dependent on anything that implements InterfaceB)
	graph.Define(&a, inject.NewAutoProvider(pkgA.NewA))

	// define how to construct b (no dependencies)
	graph.Define(&b, inject.NewProvider(pkgB.NewB))

	// resolve a and its dependency b (because *pkgB.StructB is assignable to InterfaceB)
	graph.Resolve(&a)

    // a is now usable
	a.DoStuff()
}
```

You CAN resolve everything in the graph, using `graph.ResolveAll()`, but you can also share a graph between multiple
code paths (like a controller with multiple endpoints, or a command with multiple sub-commands) and only resolve the
dependencies you need using `graph.Resolve(&ptr)`.

Because the definitions are uniquely keyed by pointer, you can also share code that produces a general graph, and
override individual definitions with more specific providers (like tests that replace a few concrete impls with mocks).

# Installation

To install Inject, use go get:

```
go get github.com/karlkfi/inject
```

# Updating

To update Inject, use go get -u:

```
go get -u github.com/karlkfi/inject
```

# Dependencies
Inject has no runtime dependencies. Tests depend on [Gomega](https://github.com/onsi/gomega).

# Testing
Tests depend on  [Gomega](https://github.com/onsi/gomega).

To install Gomega, use go get:

```
go get github.com/onsi/gomega
```

To run Inject tests, use go test:

```
go test github.com/karlkfi/inject
```

# License

   Copyright 2015 Karl Isenberg

   Licensed under the [Apache License Version 2.0](LICENSE) (the "License");
   you may not use this project except in compliance with the License.

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
