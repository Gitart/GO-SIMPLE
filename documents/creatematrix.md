# Matrix

Method of task description, library go.matrix[edit]   
Translation of: Common Lisp     
A fairly close port of the Common Lisp solution, this solution uses the go.matrix    
library for supporting functions. Note though, that go.matrix has QR decomposition,    
as shown in the Go solution to Polynomial regression. The solution there is coded more directly than    
by following the CL example here. Similarly, examination of the go.matrix QR source shows that it computes    
the decomposition more directly.    

```golang
package main
 
import (
    "fmt"
    "math"
 
    "github.com/skelterjohn/go.matrix"
)
 
func sign(s float64) float64 {
    if s > 0 {
        return 1
    } else if s < 0 {
        return -1
    }
    return 0
}
 
func unitVector(n int) *matrix.DenseMatrix {
    vec := matrix.Zeros(n, 1)
    vec.Set(0, 0, 1)
    return vec
}
 
func householder(a *matrix.DenseMatrix) *matrix.DenseMatrix {
    m := a.Rows()
    s := sign(a.Get(0, 0))
    e := unitVector(m)
    u := matrix.Sum(a, matrix.Scaled(e, a.TwoNorm()*s))
    v := matrix.Scaled(u, 1/u.Get(0, 0))
    // (error checking skipped in this solution)
    prod, _ := v.Transpose().TimesDense(v)
    β := 2 / prod.Get(0, 0)
 
    prod, _ = v.TimesDense(v.Transpose())
    return matrix.Difference(matrix.Eye(m), matrix.Scaled(prod, β))
}
 
func qr(a *matrix.DenseMatrix) (q, r *matrix.DenseMatrix) {
    m := a.Rows()
    n := a.Cols()
    q = matrix.Eye(m)
 
    last := n - 1
    if m == n {
        last--
    }
    for i := 0; i <= last; i++ {
        // (copy is only for compatibility with an older version of gomatrix)
        b := a.GetMatrix(i, i, m-i, n-i).Copy()
        x := b.GetColVector(0)
        h := matrix.Eye(m)
        h.SetMatrix(i, i, householder(x))
        q, _ = q.TimesDense(h)
        a, _ = h.TimesDense(a)
    }
    return q, a
}
```

```golang
func main() {
    // task 1: show qr decomp of wp example
    a := matrix.MakeDenseMatrixStacked([][]float64{
        {12, -51, 4},
        {6, 167, -68},
        {-4, 24, -41}})
    q, r := qr(a)
    fmt.Println("q:\n", q)
    fmt.Println("r:\n", r)
 
    // task 2: use qr decomp for polynomial regression example
    x := matrix.MakeDenseMatrixStacked([][]float64{
        {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}})
    y := matrix.MakeDenseMatrixStacked([][]float64{
        {1, 6, 17, 34, 57, 86, 121, 162, 209, 262, 321}})
    fmt.Println("\npolyfit:\n", polyfit(x, y, 2))
}
 
func polyfit(x, y *matrix.DenseMatrix, n int) *matrix.DenseMatrix {
    m := x.Cols()
    a := matrix.Zeros(m, n+1)
    for i := 0; i < m; i++ {
        for j := 0; j <= n; j++ {
            a.Set(i, j, math.Pow(x.Get(0, i), float64(j)))
        }
    }
    return lsqr(a, y.Transpose())
}
 
func lsqr(a, b *matrix.DenseMatrix) *matrix.DenseMatrix {
    q, r := qr(a)
    n := r.Cols()
    prod, _ := q.Transpose().TimesDense(b)
    return solveUT(r.GetMatrix(0, 0, n, n), prod.GetMatrix(0, 0, n, 1))
}
 
func solveUT(r, b *matrix.DenseMatrix) *matrix.DenseMatrix {
    n := r.Cols()
    x := matrix.Zeros(n, 1)
    for k := n - 1; k >= 0; k-- {
        sum := 0.
        for j := k + 1; j < n; j++ {
            sum += r.Get(k, j) * x.Get(j, 0)
        }
        x.Set(k, 0, (b.Get(k, 0)-sum)/r.Get(k, k))
    }
    return x
}
```

Output:

```
q:
 {-0.857143,  0.394286,  0.331429,
 -0.428571, -0.902857, -0.034286,
  0.285714, -0.171429,  0.942857}
r:
 { -14,  -21,   14,
    0, -175,   70,
    0,    0,  -35}

polyfit:
 {1,
 2,
 3}
Library QR, gonum/matrix[edit]
```

```golang
package main
 
import (
    "fmt"
 
    "github.com/gonum/matrix/mat64"
)
 
func main() {
    // task 1: show qr decomp of wp example
    a := mat64.NewDense(3, 3, []float64{
        12, -51, 4,
        6, 167, -68,
        -4, 24, -41,
    })
    var qr mat64.QR
    qr.Factorize(a)
    var q, r mat64.Dense
    q.QFromQR(&qr)
    r.RFromQR(&qr)
    fmt.Printf("q: %.3f\n\n", mat64.Formatted(&q, mat64.Prefix("   ")))
    fmt.Printf("r: %.3f\n\n", mat64.Formatted(&r, mat64.Prefix("   ")))
 
    // task 2: use qr decomp for polynomial regression example
    x := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    y := []float64{1, 6, 17, 34, 57, 86, 121, 162, 209, 262, 321}
    a = Vandermonde(x, 2)
    b := mat64.NewDense(11, 1, y)
    qr.Factorize(a)
    var f mat64.Dense
    f.SolveQR(&qr, false, b)
    fmt.Printf("polyfit: %.3f\n",
        mat64.Formatted(&f, mat64.Prefix("         ")))
}
 
func Vandermonde(a []float64, degree int) *mat64.Dense {
    x := mat64.NewDense(len(a), degree+1, nil)
    for i := range a {
        for j, p := 0, 1.; j <= degree; j, p = j+1, p*a[i] {
            x.Set(i, j, p)
        }
    }
    return x
}
```

Output:
```
q: ⎡-0.857   0.394   0.331⎤
   ⎢-0.429  -0.903  -0.034⎥
   ⎣ 0.286  -0.171   0.943⎦

r: ⎡ -14.000   -21.000    14.000⎤
   ⎢   0.000  -175.000    70.000⎥
   ⎣   0.000     0.000   -35.000⎦

polyfit: ⎡1.000⎤
         ⎢2.000⎥
         ⎣3.000⎦
```
