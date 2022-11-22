
```go
package main

import (
	"fmt"
	"math"
)

const Delta = 0.0001

func isConverged(d float64) bool {
	if d < 0.0 {
		d = -d
	}
	if d < Delta {
		return true
	}
	return false
}

func Sqrt(x float64) float64 {
	z := 1.0
	tmp := 0.0
	for {
		tmp = z - (z * z - x) / 2 * z
		if d := tmp - z; isConverged(d) {
			return tmp
		}
		z = tmp
	}
	return z
}

func main() {
	attempt := Sqrt(2)
	expected := math.Sqrt(2)
	fmt.Printf("attempt = %g (expected = %g) error = %g\n",
		attempt, expected, attempt - expected)
}
```

```go
package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    t, z := 0., 1.
    for {
        z, t = z - (z*z-x)/(2*z), z
        if math.Abs(t-z) < 1e-8 {
            break
        }
    }
    return z
}

func main() {
    i := 169.
    fmt.Println(Sqrt(i))
    fmt.Println(Sqrt(i) == math.Sqrt(i))
}
```

## @mpolakovic
```go
package main

func Sqrt(x float64) float64 {
    prev, z := float64(0), float64(1)
    for abs(prev - z) > 1e-8 {
        prev = z
        z = z - (z * z - x) / (2 * z)
    }
    return z
}

func abs(number float64) float64 {
    if number < 0 {
        return number * -1
    } else {
        return number
    }
}
```

## @ur0
Here _go_es another one.

```go
package main

import (
  "fmt"
  "math"
)

const Delta = 1e-3

func sqrt(x float64) float64 {
  // x is the input, float64
  // g is the guess, float64
  g := float64(1)
  for {
    // t, what we're looking at
    t := g - (g*g - x)/2 * g
    // d is the delta
    if d := math.Abs(g - t); d < Delta {
      return  t
      break
    }
    g = t
  }
  return g
}


func main() {
  i := 2
  guess := sqrt(2)
  actual := math.Sqrt(2)
  fmt.Printf("Number: %v, Guess: %g, Actual: %g, Delta: %g", i, guess, actual, math.Abs(guess - actual))
}
```

## @cahitbeyaz

Solution for 10 iterations:

```go
package main

import ("fmt"
		"math"
		)

func Sqrt(x float64) float64 {
	var zN, zNp1 float64
	zN = 1
	for i:=0;i<10;i++{
		zNp1 = zN - ((zN*zN - x) / (2 * zN))
		zN=zNp1
	}
	return float64(zNp1)
}

func main() {
	nmr:=float64(2)
	guess:=Sqrt(nmr)
	actual:=math.Sqrt(nmr)
	delta:=math.Abs(guess-actual)
	fmt.Println("Number Guess:  Actual   Delta " ,nmr ,guess ,actual,delta)
}
```

## @R4meau

```go
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	res := 1.0
	for n := 0; n < 10; n++ {
		res = res - ((res*res - x) / (2 * res))
	}
	return res
}

func main() {
	z := 169.
	fmt.Println(Sqrt(z))
	fmt.Println(math.Sqrt(z))
}
```

## @acesaif
this is nice algorithm to find sqrt programmatically ... Before I didn't know this ...

```go
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 1; i <= 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	p := 6.0
	fmt.Println(Sqrt(p)) // newton's technique  
	fmt.Println(math.Sqrt(p)) // validation by an inbuilt function 
}
```

## @tufank
```go
 package main

import (
	"fmt"
)

const diff = 1e-6

func Sqrt(x float64) float64 {
	z := x
	var oldz = 0.0
	for {
		if v := z - oldz; -diff < v && v < diff {
			return z
		} else {
			oldz = z
			z -= (z*z - x) / (2 * z)
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
```

## @donbr
I used the statement to calculate 'z' provided in the tutorial. Works very well.

```go
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Sqrt approximation of %v attempt %v = %v\n", x, i, z)
	}
	return z
}

func main() {
	x := 1525.0
	fmt.Printf("\n*** Newton Sqrt approximation for %v = %v\n", x, Sqrt(x))
	fmt.Printf("*** math.Sqrt response for %v = %v\n", x, math.Sqrt(x))
}
```

## @ceaksan
```go
 package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z = z - (z * z-x) / (2 * z)
	}
	return z
}

func main() {
	i := float64(169)
	fmt.Println(Sqrt(i), Sqrt(i) == math.Sqrt(i))
}
```

## @danong
danong commented on Mar 28

```go
import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for prev := 0.0; math.Abs(prev-z) > 1e-8; {
		prev = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(8))
}
```

## @DersioK
Simplest way to implement the code is as follows:

```go
import "fmt"

func Sqrt(x float64) float64 {
z := float64(1)
for i:=0;i<10;i++{
z -= (zz - x) / (2z)
fmt.Println(z) //Just to view each iteration of z
}
return z
}

func main() {
fmt.Println(Sqrt(2))
}
```


##  @KLVTZ


```go
package main

import (
	"fmt"
	"math"
)

func Sqrt(start float64) float64 {
	for guess := 1.0; ; {
		state := guess
		guess -= (guess*guess - start) / (2 * guess)

		if float32(state) == float32(guess) {
			return guess
		}
	}
}

func main() {
	fmt.Println(Sqrt(6.0) == math.Sqrt(6.0)) // true
	fmt.Println(Sqrt(2) == math.Sqrt(2))     // true
	fmt.Println(Sqrt(866) == math.Sqrt(866)) // true
}
```
