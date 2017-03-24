# Reflection

В гоу есть механизм, который позволяет исследовать обьекты во время выполнения программы. Он предоставляет методы исследования этих обьектов, ничего не зная об их типах. Называется он рефлексией (reflection). 
В качестве примера можно привести функцию fmt.Fprintf, которая выводит произвольный тип. 
В гоу есть стандартный механизм для определения типов - type switch:

```golang
 
     switch x := x.(type) {
       case string:
 	return x
       case int:
 	return strconv.Itoa(x)
       ...
     }  
```

Но этот свитч не подходит для более сложных типов. 
В пакете reflect есть два важных типа - Type и Value. Type - это интерфейсный тип, который имеет большое количество методов, позволяющих инспектировать обьекты, такие, как поля структур или параметры функций. Метод reflect.TypeOf в качестве параметра принимает любой интерфейс и возвращает reflect.Type. 
Функция fmt.Printf имеет специальный шаблон %T который возвращает reflect.TypeOf обьекта. 

```golang
 t := reflect.TypeOf(3) 
 fmt.Println(t) // int
 
 v := reflect.ValueOf(3)
 x := v.Interface()
 i := x.(int)
 fmt.Printf("%T %d\n", i, i) // int 3
 ```
 
В следующем примере для инспектирования обьектов мы вместо type switch будем использовать метод reflect.Value. В гоу все типы можно разбить на несколько групп: 

```
Базовые типы - bool, string, числовые типы     
Агрегатные типы - массивы и структуры     
Ссылочные типы - каналы, функции, указатели, слайсы и словари.     
```

В отдельную группу можно выделить Invalid, т.е. не имеющий типа.    

В примере создаются простые типы:
```golang
 package main
 
 import (
 	"fmt"
 	"reflect"
 	"strconv"
 	"time"
 )
 
 func Any(value interface{}) string {
 	return formatAtom(reflect.ValueOf(value))
 }
 
 func formatAtom(v reflect.Value) string {
 	switch v.Kind() {
 	case reflect.Invalid:
 		return "invalid"
 	case reflect.Int, reflect.Int8, reflect.Int16,
 		reflect.Int32, reflect.Int64:
 		return strconv.FormatInt(v.Int(), 10)
 	case reflect.Uint, reflect.Uint8, reflect.Uint16,
 		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
 		return strconv.FormatUint(v.Uint(), 10)
 	case reflect.Bool:
 		return strconv.FormatBool(v.Bool())
 	case reflect.String:
 		return strconv.Quote(v.String())
 	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
 		return v.Type().String() + " 0x" +
 			strconv.FormatUint(uint64(v.Pointer()), 16)
 	default: // reflect.Array, reflect.Struct, reflect.Interface
 		return v.Type().String() + " value"
 	}
 }
 
 func main() {
   
 	var x int64 = 1
 	var d time.Duration = 1 * time.Nanosecond
 	fmt.Println(Any(x))                  // "1"
 	fmt.Println(Any(d))                  // "1"
 	fmt.Println(Any([]int64{x}))         // "[]int64 0x8202b87b0"
 	fmt.Println(Any([]time.Duration{d})) // "[]time.Duration 0x8202b87e0"
 }  
```

Этот пример не подходит для композитных типов вроде структур. В следующем примере создается обьект структуры, после чего выводятся ее поля:

```golang
 package main
 
 import (
 	"fmt"
 	"reflect"
 	"strconv"
 )
 
 func Display(name string, x interface{}) {
 	fmt.Printf("Display %s (%T):\n", name, x)
 	display(name, reflect.ValueOf(x))
 }
 
 func formatAtom(v reflect.Value) string {
 	switch v.Kind() {
 	case reflect.Invalid:
 		return "invalid"
 	case reflect.Int, reflect.Int8, reflect.Int16,
 		reflect.Int32, reflect.Int64:
 		return strconv.FormatInt(v.Int(), 10)
 	case reflect.Uint, reflect.Uint8, reflect.Uint16,
 		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
 		return strconv.FormatUint(v.Uint(), 10)
 	// ...floating-point and complex cases omitted for brevity...
 	case reflect.Bool:
 		if v.Bool() {
 			return "true"
 		}
 		return "false"
 	case reflect.String:
 		return strconv.Quote(v.String())
 	case reflect.Chan, reflect.Func, reflect.Ptr,
 		reflect.Slice, reflect.Map:
 		return v.Type().String() + " 0x" +
 			strconv.FormatUint(uint64(v.Pointer()), 16)
 	default: // reflect.Array, reflect.Struct, reflect.Interface
 		return v.Type().String() + " value"
 	}
 }
 
 func display(path string, v reflect.Value) {
 	switch v.Kind() {
 	case reflect.Invalid:
 		fmt.Printf("%s = invalid\n", path)
 	case reflect.Slice, reflect.Array:
 		for i := 0; i < v.Len(); i++ {
 			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
 		}
 	case reflect.Struct:
 		for i := 0; i < v.NumField(); i++ {
 			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
 			display(fieldPath, v.Field(i))
 		}
 	case reflect.Map:
 		for _, key := range v.MapKeys() {
 			display(fmt.Sprintf("%s[%s]", path,
 				formatAtom(key)), v.MapIndex(key))
 		}
 	case reflect.Ptr:
 		if v.IsNil() {
 			fmt.Printf("%s = nil\n", path)
 		} else {
 			display(fmt.Sprintf("(*%s)", path), v.Elem())
 		}
 	case reflect.Interface:
 		if v.IsNil() {
 			fmt.Printf("%s = nil\n", path)
 		} else {
 			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
 			display(path+".value", v.Elem())
 		}
 	default: // basic types, channels, funcs
 		fmt.Printf("%s = %s\n", path, formatAtom(v))
 	}
 }
 ```
 
 ### Sample 2
 
 ```golang
  
 func main() {
 
 type Movie struct {
   Title, Subtitle string
   Year int
   Color bool
   Actor map[string]string
   Oscars []string
   Sequel *string
 }
 	
 	
 movie := Movie{
   Title: "title",
   Subtitle: "subtitle",
   Year: 1964,
   Color: false,
   Actor: map[string]string{
     "Brig. Gen. Jack D. Ripper": "SterlingHayden",
     `Maj. T.J. "King" Kong`: "Slim Pickens",
   },
   Oscars: []string{
     "Best Director (Nomin.)",
     "Best Picture (Nomin.)",
   },
 }
 
  Display("movie", movie)
  
 }
 ```
 
 Вывод:
 ```
 Display movie (main.Movie):
 movie.Title = "title"
 movie.Subtitle = "subtitle"
 movie.Year = 1964
 movie.Color = false
 movie.Actor["Brig. Gen. Jack D. Ripper"] = "SterlingHayden"
 movie.Actor["Maj. T.J. \"King\" Kong"] = "Slim Pickens"
 movie.Oscars[0] = "Best Director (Nomin.)"
 movie.Oscars[1] = "Best Picture (Nomin.)"
 movie.Sequel = nil
 ```
 
Если у обьекта есть методы, их список можно распечатать, для этого можно использовать reflect.Type. В следующем примере распечатываются методы стандартных обьектов.

 ```golang
 package main
 
 import (
 	"fmt"
 	"reflect"
 	"strings"
 	"time"
 	
 )
 
 func Print(x interface{}) {
 	v := reflect.ValueOf(x)
 	t := v.Type()
 	fmt.Printf("type %s\n", t)
 
 	for i := 0; i < v.NumMethod(); i++ {
 		methType := v.Method(i).Type()
 		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name,
 			strings.TrimPrefix(methType.String(), "func"))
 	}
 }
 
 func main() {
  
   Print(time.Hour)
   Print(new(strings.Replacer))
 }  
 ```
 
 Вывод:
 ```
 type time.Duration
 func (time.Duration) Hours() float64
 func (time.Duration) Minutes() float64
 func (time.Duration) Nanoseconds() int64
 func (time.Duration) Seconds() float64
 func (time.Duration) String() string
 type *strings.Replacer
 func (*strings.Replacer) Replace(string) string
 func (*strings.Replacer) WriteString(io.Writer, string) (int, error)
 ```
 
