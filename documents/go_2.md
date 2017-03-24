# Программирование в Go

### Часть 1

В go указатели поддерживаются ограниченно - в частности, не поддерживается арифметика указателей. В go нет free() или delete(), но есть гарбаж-коллектор. В go переменные как правило хранят значения, но могут хранить и ссылки - как в случае с каналами, функциями, методами, мапами и слайсами. Значения, передаваемые в функции и методы, копируются, т.е. создаются заново - это справедливо для чисел и строк, которые immutable. Это справедливо также для массивов и структур. Слайсы и мапы же передаются по ссылке, поскольку являются референсными типами. Указатель в go - это переменная, которая хранит адрес другой переменной. & - это ссылка, равная этому адресу памяти. Следующий рисунок иллюстрирует вышесказанное: 
В go указатель может указывать на другой указатель. 

Пример    

```golang
 z := 37      // z is of type int
 pi := &z     // pi is of type *int (pointer to int)
 ppi := &pi   // ppi is of type **int (pointer to pointer to int)
 fmt.Println(z, *pi, **ppi)
 **ppi++      // то же самое, что: (*(*ppi))++ или: *(*ppi)++
 fmt.Println(z, *pi, **ppi)
 ```
 
 Вывод:
 ```
 37 37 37
 38 38 38
 ```
 
Еще пример:

```golang
 i := 9
 j := 5
 product := 0
 swapAndProduct1(&i, &j, &product)
 fmt.Println(i, j, product)
 
 func swapAndProduct1(x, y, product *int) {
   if *x > *y {
   *x, *y = *y, *x
   }
   *product = *x * *y // The compiler would be }
 }
```

 Вывод:
 ```
 5 9 45
 ```
 
Последняя функция может быть переписана более понятным языком, но у нее будет минус - она своп делает не по месту, а создает для этого локальные переменные:
```golang
 func swapAndProduct2(x, y int) (int, int, int) {
   if x > y {
   x, y = y, x
   }
   return x, y, x * y
 }
 ```
 
Для того чтобы в go создать переменную-указатель и присвоить ей значение, есть два варианта - обычный стандартный и с помощью оператора new() . Рассмотрим пример со структурой:

```golang
 type composer struct {
   name string
   birthYear int
 }
 
 antónio := composer{"António Teixeira", 1707}	         // composer value
 agnes := new(composer) 				 // pointer to composer
 agnes.name, agnes.birthYear = "Agnes Zimmermann", 1845
 julia := &composer{}  				         // pointer to composer
 julia.name, julia.birthYear = "Julia Ward Howe", 1819
 augusta := &composer{"Augusta Holmès", 1847}            // pointer to composer
 fmt.Println(antónio)
 fmt.Println(agnes, augusta, julia)
```

 Вывод:
```
{António Teixeira 1707}
 &{Agnes Zimmermann 1845} &{Augusta Holmès 1847} &{Julia Ward Howe 1819}
 ```
 
Следующий пример показывает, что когда в функцию передаются референсные типы, такие, как слайсы и мапы, все изменения, которые с ними происходят внутри функции, будут видны снаружи:

```golang
 grades := []int{87, 55, 43, 71, 60, 43, 32, 19, 63}
 inflate(grades, 3)
 fmt.Println(grades)
 
 func inflate(numbers []int, factor int) {
   for i := range numbers {
     numbers[i] *= factor
   }
 }
``` 
 Вывод:
 ```
 [261 165 129 213 180 129 96 57 189]
 ```
 
Если же мы передадим в функцию структуру, изменения внутри функции снаружи видны не будут, поскольку структура - это value-тип. Но если мы передадим в функцию ссылку на структуру, то будут:

```golang
 type rectangle struct {
   x0, y0, x1, y1 int
   fill  color.RGBA
 }
 
 rect := rectangle{4, 8, 20, 10, color.RGBA{0xFF, 0, 0, 0xFF}}
 fmt.Println(rect)
 resizeRect(&rect, 5, 5)
 fmt.Println(rect)
 
 func resizeRect(rect *rectangle, Δwidth, Δheight int) {
   (*rect).x1 += Δwidth    // Ugly explicit dereference
   rect.y1 += Δheight      // . automatically dereferences structs
 }
```

 Вывод:
```
 {4 8 20 10 {255 0 0 255}}
 {4 8 25 15 {255 0 0 255}}
 ```
 
Массив в go - это коллекция фиксированного размера. Пример:

```golang
 var buffer [20]byte
 var grid1 [3][3]int
 grid1[1][0], grid1[1][1], grid1[1][2] = 8, 6, 2
 grid2 := [3][3]int{{4, 3}, {8, 6, 2}}
 cities := [...]string{"Shanghai", "Mumbai", "Istanbul", "Beijing"}
 cities[len(cities)-1] = "Karachi"
 fmt.Println("Type Len Contents")
 fmt.Printf("%-8T %2d %v\n", buffer, len(buffer), buffer)
 fmt.Printf("%-8T %2d %q\n", cities, len(cities), cities)
 fmt.Printf("%-8T %2d %v\n", grid1, len(grid1), grid1)
 fmt.Printf("%-8T %2d %v\n", grid2, len(grid2), grid2)
```

 Вывод:
```
 Type           Len    Contents
 [20]uint8      20     [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
 [4]string      4      ["Shanghai" "Mumbai" "Istanbul" "Karachi"]
 [3][3]int      3      [[0 0 0] [8 6 2] [0 0 0]]
 [3][3]int      3      [[4 3 0] [8 6 2] [0 0 0]]
 ```
 
Слайс в go - это коллекция переменного размера. У него есть функция append(). Хранимые обьекты по умолчанию должны быть одного типа, но нам никто не мешает создать массив или слайс пустых интерфейсов, и можно будет хранить обьекты произвольного типа, Проблемы будут на этапе извлечения элементов, Слайсы можно создавать двумя способами - стандартным и встроенной функцией make(), с помощью которой кстати также можно создавать мапы и каналы. У слайса есть две встроенных функции - len() и cap(). cap() отличается от len() тем, что равен максимально возможному числу элементов. Пример:

 s := []string{"A", "B", "C", "D", "E", "F", "G"}
 t := s[2:6]
 fmt.Println(t, s, "=", s[:4], "+", s[4:])
 s[3] = "x"
 t[len(t)-1] = "y"
 fmt.Println(t, s, "=", s[:4], "+", s[4:])
 
 Вывод:
 [C D E F] [A B C D E F G] = [A B C D] + [E F G]
 [C x E y] [A B C x E y G] = [A B C x] + [E y G]
Пример:

 buffer := make([]byte, 20, 60)
 grid1 := make([][]int, 3)
 for i := range grid1 {
     grid1[i] = make([]int, 3)
 }
 grid1[1][0], grid1[1][1], grid1[1][2] = 8, 6, 2
 grid2 := [][]int{{4, 3, 0}, {8, 6, 2}, {0, 0, 0}}
 cities := []string{"Shanghai", "Mumbai", "Istanbul", "Beijing"}
 cities[len(cities)-1] = "Karachi"
 fmt.Println("Type Len Cap Contents")
 fmt.Printf("%-8T %2d %3d %v\n", buffer, len(buffer), cap(buffer), buffer)
 fmt.Printf("%-8T %2d %3d %q\n", cities, len(cities), cap(cities), cities)
 
 Вывод:
 Type		Len	Cap	 Contents
 []uint8 	20 	60 	[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
 []string 	4	4       ["Shanghai" "Mumbai" "Istanbul" "Karachi"]
Для итерации слайсов используется цикл for ... range:

 amounts := []float64{237.81, 261.87, 273.93, 279.99, 281.07, 303.17,231.47, 227.33, 209.23, 197.09}
 sum := 0.0
 for _, amount := range amounts {
   sum += amount
 }
 fmt.Printf("Σ %.1f → %.1f\n", amounts, sum)
 
 Вывод:
 Σ [237.8 261.9 273.9 280.0 281.1 303.2 231.5 227.3 209.2 197.1] → 2503.0
Если мы хотим модифицировать сам слайс прямо во время его итерации:

 for i := range amounts {
   amounts[i] *= 1.05
   sum += amounts[i]
 }
Итерация слайса, состоящего из структур - здесь возможна модификация слайса, поскольку итерируется указатель на слайс:

 type Product struct {
   name string
   price float64
 }
 
 func (product Product) String() string {
   return fmt.Sprintf("%s (%.2f)", product.name, product.price)
 }
 
 products := []*Product{{"Spanner", 3.99}, {"Wrench", 2.49},{"Screwdriver", 1.99}}
 fmt.Println(products)
 for _, product := range products {
   product.price += 0.50
 }
 fmt.Println(products)
 
 Вывод:
 [Spanner (3.99) Wrench (2.49) Screwdriver (1.99)]
 [Spanner (4.49) Wrench (2.99) Screwdriver (2.49)]
Функция append() добавляет новый элемент в слайс:

 s := []string{"A", "B", "C", "D", "E", "F", "G"}
 t := []string{"K", "L", "M", "N"}
 u := []string{"m", "n", "o", "p", "q", "r"}
 s = append(s, "h", "i", "j") // Append individual values
 s = append(s, t...) // Append all of a slice's values
 s = append(s, u[2:5]...) // Append a subslice
 b := []byte{'U', 'V'}
 letters := "wxy"
 b = append(b, letters...) // Append a string's bytes to a byte slice
 fmt.Printf("%v\n%s\n", s, b)
 
 Вывод:
 [A B C D E F G h i j K L M N o p q]
 UVwxy
Для того, чтобы вставить элемент в произвольное место слайса, прийдется писать функцию:

 s := []string{"M", "N", "O", "P", "Q", "R"}
 x := InsertStringSliceCopy(s, []string{"a", "b", "c"}, 0) // At the front
 y := InsertStringSliceCopy(s, []string{"x", "y"}, 3) // In the middle
 z := InsertStringSliceCopy(s, []string{"z"}, len(s)) // At the end
 fmt.Printf("%v\n%v\n%v\n%v\n", s, x, y, z)
 
 func InsertStringSliceCopy(slice, insertion []string, index int) []string {
   result := make([]string, len(slice)+len(insertion))
   at := copy(result, slice[:index])
   at += copy(result[at:], insertion)
   copy(result[at:], slice[index:])
   return result
 }
 
 Вывод:
 
 [M N O P Q R]
 [a b c M N O P Q R]
 [M N O x y P Q R]
 [M N O P Q R z]
Функция вставки может иметь следующий вид - ее отличие от предыдущего варианта в том, что она меняет существующий слайс, в то время как предыдущая создает новый:

 func InsertStringSlice(slice, insertion []string, index int) []string {
   return append(slice[:index], append(insertion, slice[index:]...)...)
 }
Можно удалять элементы из слайса с произвольной позиции:

 s := []string{"A", "B", "C", "D", "E", "F", "G"}
 s = s[2:] // Remove s[:2] from the front
 fmt.Println(s)
 
 Вывод:
 [C D E F G]
 
 s := []string{"A", "B", "C", "D", "E", "F", "G"}
 s = s[:4] // Remove s[4:] from the end
 fmt.Println(s)
 
 Вывод:
 [A B C D]
 
 s := []string{"A", "B", "C", "D", "E", "F", "G"}
 s = append(s[:1], s[5:]...) // Remove s[1:5] from the middle
 fmt.Println(s)
 
 Вывод:
 [A F G]
Стандартная библиотека содержит функции сортировки слайсов, состоящих из чисел или строк.

 files := []string{"Test.conf", "util.go", "Makefile", "misc.go", "main.go"}
 fmt.Printf("Unsorted: %q\n", files)
 sort.Strings(files) // Standard library sort function
 fmt.Printf("Underlying bytes: %q\n", files)
 SortFoldedStrings(files) // Custom sort function
 fmt.Printf("Case insensitive: %q\n", files)
 
 func SortFoldedStrings(slice []string) {
   sort.Sort(FoldedStrings(slice))
 }
 
 type FoldedStrings []string
 
 func (slice FoldedStrings) Len() int { return len(slice) }
 
 func (slice FoldedStrings) Less(i, j int) bool {
   return strings.ToLower(slice[i]) < strings.ToLower(slice[j])
 }
 
 func (slice FoldedStrings) Swap(i, j int) {
   slice[i], slice[j] = slice[j], slice[i]
 }
 
 Вывод:
 Unsorted:["Test.conf" "util.go" "Makefile" "misc.go" "main.go"]
 Underlying bytes: ["Makefile" "Test.conf" "main.go" "misc.go" "util.go"]
 Case insensitive: ["main.go" "Makefile" "misc.go" "Test.conf" "util.go"]
Найти в слайсе индекс по значению можно без всяких функций:

 files := []string{"Test.conf", "util.go", "Makefile", "misc.go", "main.go"}
 target := "Makefile"
 for i, file := range files {
   if file == target {
     fmt.Printf("found \"%s\" at files[%d]\n", file, i)
     break
   }
 }
 
 Вывод:
 found "Makefile" at files[2]
С использованием стандартных поисковых функций:

 sort.Strings(files)
 fmt.Printf("%q\n", files)
 i := sort.Search(len(files),
   func(i int) bool { return files[i] >= target })
 if i < len(files) && files[i] == target {
   fmt.Printf("found \"%s\" at files[%d]\n", files[i], i)
 }
 
 Вывод:
 ["Makefile" "Test.conf" "main.go" "misc.go" "util.go"]
 found "Makefile" at files[0]
Map - неотсортированная коллекция пар ключ-значение не-фиксированного размера. В качестве ключа используются встроенные типы. Мапы, или словари - это ссылочный тип. Поиск по ключу в мапах шустрее линейного поиска. Ключи в мапах должны быть одного типа, равно как и значения. Операции над мапами:

 m[k] = v
 delete(m, k)
 v := m[k]
 v, found := m[k] 
 len(m)
Мапы создаются с помощью функции make:

 massForPlanet := make(map[string]float64) // Same as: map[string]float64{}
 massForPlanet["Mercury"] = 0.06
 massForPlanet["Venus"] = 0.82
 massForPlanet["Earth"] = 1.00
 massForPlanet["Mars"] = 0.11
 fmt.Println(massForPlanet)
 
 Dsdjl^
 map[Venus:0.82 Mars:0.11 Earth:1 Mercury:0.06]
В качестве ключа в мапах также можно использовать указатели:

 type Point struct{ x, y, z int }
 
 func (point Point) String() string {
   return fmt.Sprintf("(%d,%d,%d)", point.x, point.y, point.z)
 }
 
 triangle := make(map[*Point]string, 3)
 triangle[&Point{89, 47, 27}] = "α"
 triangle[&Point{86, 65, 86}] = "β"
 triangle[&Point{7, 44, 45}] = "γ"
 fmt.Println(triangle)
 
 Вывод:
 map[(7,44,45):γ (89,47,27):α (86,65,86):β]
Более того, в go позволительно в качестве ключа использовать структуру:

 nameForPoint := make(map[Point]string) // Same as: map[Point]string{}
 nameForPoint[Point{54, 91, 78}] = "x"
 nameForPoint[Point{54, 158, 89}] = "y"
 fmt.Println(nameForPoint)
 
 Вывод
 map[(54,91,78):x (54,158,89):y]
Итерация мапов:

 populationForCity := map[string]int{"Istanbul": 12610000,
     "Karachi": 10620000, "Mumbai": 12690000, "Shanghai": 13680000}
 for city, population := range populationForCity {
     fmt.Printf("%-10s %8d\n", city, population)
 }
 
 Вывод:
 Shanghai        13680000
 Mumbai          12690000
 Istanbul        12610000
 Karachi         10620000
Для того чтобы вывести значения мапа в алфавитном порядке, нужно скопировать мап в слайс, потом отсортировать слайс:

 cities := make([]string, 0, len(populationForCity))
 for city := range populationForCity {
   cities = append(cities, city)
 }
 
 sort.Strings(cities)
 for _, city := range cities {
   fmt.Printf("%-10s %8d\n", city, populationForCity[city])
 }
 
 Вывод:
 Beijing        11290000
 Istanbul       12610000
 Karachi        11620000
 Mumbai         12690000
