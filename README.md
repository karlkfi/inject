# Inject
Dependency injection library for Go (golang)

# Why use Inject?

Unlike most other golang injection frameworks, Inject uses constructor functions to resolve nodes and relationships in the dependency graph.
This has several advantages: 
- Public access to the struct is not required.
- No coordination is required between implementations of a common interface.
- No explicit name is required for nodes in the dependency graph. The name is implicit, using a pointer.

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
		primative = "some string"
		a    pkgA.InterfaceA
		b    *pkgB.StructB
	)

	graph.Define(&a, inject.NewProvider(pkgA.NewA, &b))
	graph.Define(&b, inject.NewProvider(pkgB.NewB, &primative))
	graph.ResolveAll()

	a.DoStuff()
}

```

In the simple example above, the user defines two dependency relationships: 
- a depends on b (using constructor pkgA.NewA)
- b depends on a primative (using constructor pkgA.NewB)

# Go Docs

Generated documentation is available on [godoc.org](https://godoc.org/github.com/karlkfi/inject)

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
Inject has no runtime dependencies. Tests depend on [testify](https://github.com/stretchr/testify). 

# Testing
Tests depend on [testify](https://github.com/stretchr/testify). 

To install Testify, use go get:

```
go get github.com/stretchr/testify
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
