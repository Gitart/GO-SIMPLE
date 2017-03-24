# Базовые типы в Go


## Идентификатор в go - это набор символов, букв, символа подчеркивания. Зарезервированные ключевые слова в go:
```
 break    default     func     interface     select
 case     defer       go       map           struct
 chan     else        goto     package       switch
 const    fallthrough if       range         type
 continue for         import   return        var
 ```
 
## Зарезервированные идентификаторы:
 ```
 append         copy     int8       nil      true
 bool           delete   int16      panic    uint
 byte           error    int32      print    uint8
 cap            false    int64      println  uint16
 close          float32  iota       real     uint32
 complex        float64  len        recover  uint64
 complex64      imag     make       rune     uintptr
 complex128     int      new        string
 ```
 
## Константы декларируются с помощью ключевого слова const. Переменные можно декларировать с помошью ключевого слова var или без:
```
 const limit = 512			// constant; type-compatible with any number
 const top uint16 = 1421                // constant; type: uint16
 start := -19				// variable; inferred type: int
 end := int64(9876543210)		// variable; type: int64
 var i int				// variable; value 0; type: int
 var debug = false			// variable; inferred type: bool
 checkResults := true			// variable; inferred type: bool
 stepSize := 1.5                        // variable; inferred type: float64
 acronym := "FOSS"			// variable; inferred type: string
 ```
 
В go два булевских типа true и false. 
Бинарные логические операторы - ||(или) и &&(и). 
Операторы сравнения - <, <=, ==, !=, >=, > 
В Go 11 целочисленных типов, 5 знаковых, 5 беззнаковых, плюс указатель:
```
 byte                 Synonym for uint8
 int                  The int32 or int64 range depending on the implementation
 int8                 [−128, 127]  
 int16                [−32768, 32767]
 int32                [−2147483648, 2147483647]
 int64                [−9223372036854775808, 9223372036854775807]
 rune                 Synonym for int32
 uint                 The uint32 or uint64 range depending on the implementation
 uint8                [0, 255]
 uint16               [0, 65535]
 uint32               [0, 4294967295]
 uint64               [0, 18446744073709551615]
 uintptr              An unsigned integer capable of storing a pointer value (advanced)
 ```
 
В Go есть два типа чисел с плавающей точкой и два типа комплексных чисел:
```
 float32 	±3.40282346638528859811704183484516925440 х 10^38
 float64        ±1.797693134862315708145274237317043567981 х 10^308
 complex64      The real and imaginary parts are both of type float32.
 complex128 	The real and imaginary parts are both of type float64.
 ```
 
### Например:

```golang
 f := 3.2e5                         // type: float64
 x := -7.3 - 8.9i                   // type: complex128 (literal)
 y := complex64(-18.3 + 8.9i)       // type: complex64 (conversion) 
 z := complex(f, 13.2)              // type: complex128 (construction)
 fmt.Println(x, real(y), imag(z))   // Prints: (-7.3-8.9i) -18.3 13.2
 ```
 
