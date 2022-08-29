# Polymorphism in Go Programming Language

---

Polymorphism is the ability to write code that can take on different behavior through the implementation of types.

We have the declaration of a structs named Pentagon, Hexagon, Octagon and Decagon with the implementation of the **Geometry** interface.

We have our polymorphic Edges functions that accepts values that implement the **Geometry** interface. Using polymorphic approach the method created here Parameter is used by each concrete type value that's passed in.

### Example

```jsx
package main

import (
	"fmt"
)

// Geometry is an interface that defines Geometrical Calculation
type Geometry interface {
	Edges() int
}

// Pentagon defines a geometrical object
type Pentagon struct{}

// Hexagon defines a geometrical object
type Hexagon struct{}

// Octagon defines a geometrical object
type Octagon struct{}

// Decagon defines a geometrical object
type Decagon struct{}

// Edges implements the Geometry interface
func (p Pentagon) Edges() int { return 5 }

// Edges implements the Geometry interface
func (h Hexagon) Edges() int { return 6 }

// Edges implements the Geometry interface
func (o Octagon) Edges() int { return 8 }

// Edges implements the Geometry interface
func (d Decagon) Edges() int { return 10 }

// Parameter calculate parameter of object
func Parameter(geo Geometry, value int) int {
	num := geo.Edges()
	calculation := num * value
	return calculation
}

// main is the entry point for the application.
func main() {
	p := new(Pentagon)
	h := new(Hexagon)
	o := new(Octagon)
	d := new(Decagon)

	g := [...]Geometry{p, h, o, d}

	for _, i := range g {
		fmt.Println(Parameter(i, 5))
	}
}

```

### Output

```jsx
25
30
40
50
```
