![image](https://user-images.githubusercontent.com/3950155/189183250-402bf8e4-9638-429e-bfe5-7521553ce631.png)


Ключевое слово Atomic позволяет нам выполнять синхронные операции. Это пакет языка go, который используется для управления синхронным поведением языка.
Please see the below syntax for understanding.

```
waitgroup  sync.WaitGroup
Waitgroup.waitgroup-function
waitgroup.Wait()
```


### Example 1
```go
package main
 
import (
    "fmt"
    "sync"
    "sync/atomic"
)
 
func f(v *uint32, wg *sync.WaitGroup) {
    for i := 0; i < 3000; i++ {
        atomic.AddUint32(v, 1)
    }
    wg.Done()
}
 
func main() {
    var v uint32 = 42
    var wg sync.WaitGroup
    wg.Add(2)
    go f(&v, &wg)
    go f(&v, &wg)
    wg.Wait()
 
    fmt.Println(v)
}
```

### Example 2
```go
package main
 
import (
    "fmt"
    "sync"
    // "sync/atomic"
)
 
func f(v *int, wg *sync.WaitGroup) {
    for i := 0; i < 3000; i++ {
        *v++
    }
    wg.Done()
}
 
func main() {
    var v int = 42
    var wg sync.WaitGroup
    wg.Add(2)
    go f(&v, &wg)
    go f(&v, &wg)
    wg.Wait()
 
    fmt.Println(v)
}
```

### Example 3
```go
package main
import (
"fmt"
)
import (
"sync/atomic"
)
import (
"sync"
)
import (
"runtime"
)
//Initialising the wait group variable for further uses
//Here we are also defining the variable val as the int32 which will be used further in the program
var (
val int32
waitgroup sync.WaitGroup
)
func main() {
//calling the Add on the wait group
waitgroup.Add(3)
//Creating the channel of the grouting for performing operations one after another
go append("PHP")
go append("Python")
go append("Go")
//performing the wait operation the variable created from the atomic sync package .
waitgroup.Wait()
fmt.Println("The value of the counter is:", val)
}
func append(lang string) {
defer waitgroup.Done()
for range lang {
atomic.AddInt32(&val, 3)
runtime.Gosched()
}
}
```

### Example 4

```go
package main
import (
"fmt"
)
import (
"sync"
)
import (
"sync/atomic"
)
func main() {
var val int32
//Initialising the wait group variable for further uses
var waitgroup sync.WaitGroup
for k := 1; k < 51; k++ {
waitgroup.Add(1)
go func() {
for l := 1; l < 1001; l++ {
atomic.AddInt32(&val, 1)
}
waitgroup.Done()
}()
}
//performing the wait operation the variable created from the atomic sync package .
waitgroup.Wait()
fmt.Println("The value of the val is:", val)
}
```
