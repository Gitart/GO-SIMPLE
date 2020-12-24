# Most Popular Golang Slice Sort, Reverse, Search Functions

Slice sorting or searching functions allow you to interact with and manipulate slice in various ways. Golang sort functions are the part of the core. There is no installation required to use this function only you need to import "sort" package. With the help of sort function you can search any A list of important Golang sort functions are as follow:

## 1) Golang *Ints* function \[ascending order\]

The Ints function sorts a slice of integer in ascending order.

###### Syntax:

func Ints(intSlice \[\]int)

###### Example:

package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := \[\]int{10, 5, 25, 351, 14, 9} // unsorted
	fmt.Println("Slice of integer BEFORE sort:",intSlice)
	sort.Ints(intSlice)
	fmt.Println("Slice of integer AFTER  sort:",intSlice)
}

###### When you run the program, you get the following output:

C:\\golang>go run sort.go Slice of integer BEFORE sort: \[10 5 25 351 14 9\] Slice of integer AFTER sort: \[5 9 10 14 25 351\] C:\\golang>

## 2) Golang *Strings* function \[ascending order\]

The Strings function sorts a slice of strings in ascending order lexicographically.

###### Syntax:

func Strings(strSlice \[\]string)

###### Example:

package main

import (
	"fmt"
	"sort"
)

func main() {
	strSlice := \[\]string{"Jamaica","Estonia","Indonesia","Hong Kong"} // unsorted
	fmt.Println("Slice of string BEFORE sort:",strSlice)
	sort.Strings(strSlice)
	fmt.Println("Slice of string AFTER  sort:",strSlice)

	fmt.Println("\\n\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\\n")

	strSlice = \[\]string{"JAMAICA","Estonia","indonesia","hong Kong"} // unsorted
	fmt.Println("Slice of string BEFORE sort:",strSlice)
	sort.Strings(strSlice)
	fmt.Println("Slice of string AFTER  sort:",strSlice)
}

###### When you run the program, you get the following output:

C:\\golang>go run sort.go Slice of string BEFORE sort: \[Jamaica Estonia Indonesia Hong Kong\] Slice of string AFTER sort: \[Estonia Hong Kong Indonesia Jamaica\] \-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\- Slice of string BEFORE sort: \[JAMAICA Estonia indonesia hong Kong\] Slice of string AFTER sort: \[Estonia JAMAICA hong Kong indonesia\] C:\\golang>

## 3) Golang *Float64s* function \[ascending order\]

The Float64s function sorts a slice of float64 in ascending order.

###### Syntax:

func Float64s(fltSlice \[\]string)

###### Example:

package main

import (
	"fmt"
	"sort"
)

func main() {
	fltSlice := \[\]float64{18787677.878716, 565435.321, 7888.545, 8787677.8716, 987654.252} // unsorted
	fmt.Println("Slice BEFORE sort: ",fltSlice)

	sort.Float64s(fltSlice)

	fmt.Println("Slice AFTER sort: ",fltSlice)
}

###### When you run the program, you get the following output:

C:\\golang>go run sort.go Slice BEFORE sort: \[1.8787677878716e+07 565435.321 7888.545 8.7876778716e+06 987654.252\] Slice AFTER sort: \[7888.545 565435.321 987654.252 8.7876778716e+06 1.8787677878716e+07\] C:\\golang>

## 4) Golang *IntsAreSorted* function

The IntsAreSorted function tests whether a slice of integer is sorted in ascending order. Returns true if the slice of numbers is found in the ascending order, or false otherwise.

###### Syntax:

func IntsAreSorted(a \[\]string) bool

###### Example:

package main

import (
    "fmt"
    "sort"
)

func main() {
    intSlice := \[\]int{10, 5, 25, 351, 14, 9}	// unsorted
	fmt.Println(sort.IntsAreSorted(intSlice))	// false

	intSlice = \[\]int{5, 9, 14, 351, 614, 999}	// sorted
	fmt.Println(sort.IntsAreSorted(intSlice))	// true
}

###### When you run the program, you get the following output:

C:\\golang>go run sort.go false true C:\\golang>

## 5) Golang *StringsAreSorted* function

The StringsAreSorted function tests whether a slice of string is sorted in ascending order. Returns true if the slice of string is found in the ascending order, or false otherwise.

###### Syntax:

func StringsAreSorted(strSlice \[\]string) bool

###### Example:

package main

import (
    "fmt"
    "sort"
)

func main() {
    strSlice := \[\]string{"Jamaica","Estonia","Indonesia","Hong Kong"} // unsorted
    fmt.Println(sort.StringsAreSorted(strSlice))	// false

    strSlice = \[\]string{"JAMAICA","Estonia","indonesia","hong Kong"} // unsorted
    fmt.Println(sort.StringsAreSorted(strSlice))	// false

	strSlice = \[\]string{"estonia","hong Kong","indonesia","jamaica"} // sorted
    fmt.Println(sort.StringsAreSorted(strSlice))	// true
}

###### When you run the program, you get the following output:

C:\\golang>go run sort.go false false true C:\\golang>

## 6) Golang *Float64sAreSorted* function

The Float64sAreSorted function tests whether a slice of float64s is sorted in ascending order. Returns true if the slice of float64 is found in the ascending order, or false otherwise.

###### Syntax:

func Float64sAreSorted(fltSlice \[\]float64) bool

###### Example:

package main

import (
	"fmt"
	"sort"
)

func main() {
	fltSlice := \[\]float64{18787677.878716, 565435.321, 7888.545, 8787677.8716, 987654.252} // unsorted
	fmt.Println(sort.Float64sAreSorted(fltSlice))	// false

	fltSlice = \[\]float64{565435.321, 887888.545, 8787677.8716, 91187654.252} // sorted
	fmt.Println(sort.Float64sAreSorted(fltSlice))	// true
}

###### When you run the program, you get the following output:

C:\\golang>go run sort.go false true C:\\golang>

## 7) Golang *SearchInts* function \[ascending order\]

The SearchInts function searches the position of x in a sorted slice of int and returns the index as specified by Search. This function works if slice is in sort order only. If found x in intSlice then it returns index position of intSlice otherwise it returns index position where x fits in sorted slice. The following example shows the usage of SearchInts() function:

###### Syntax:

func SearchInts(intSlice \[\]int, x int) int

###### Example:

package main

import (
	"fmt"
	"sort"
)

func main() {
	// integer slice in unsort order
	intSlice := \[\]int{55, 22, 18, 9, 12, 82, 28, 36, 45, 65}
	x := 18
	pos := sort.SearchInts(intSlice,x)
	fmt.Printf("Found %d at index %d in %v\\n", x, pos, intSlice)

	// slice need to be sort in ascending order before to use SearchInts
	sort.Ints(intSlice)	// slice sorted
	pos = sort.SearchInts(intSlice,x)
	fmt.Printf("Found %d at index %d in %v\\n", x, pos, intSlice)

	x = 54
	pos = sort.SearchInts(intSlice,x)
	fmt.Printf("Found %d at index %d in %v\\n", x, pos, intSlice)

	x = 99
	pos = sort.SearchInts(intSlice,x)
	fmt.Printf("Found %d at index %d in %v\\n", x, pos, intSlice)

	x = \-5
	pos = sort.SearchInts(intSlice,x)
	fmt.Printf("Found %d at index %d in %v\\n", x, pos, intSlice)
}

###### When you run the program, you get the following output:

C:\\golang>go run sort.go Found 18 at index 0 in \[55 22 18 9 12 82 28 36 45 65\] Found 18 at index 2 in \[9 12 18 22 28 36 45 55 65 82\] Found 54 at index 7 in \[9 12 18 22 28 36 45 55 65 82\] Found 99 at index 10 in \[9 12 18 22 28 36 45 55 65 82\] Found \-5 at index 0 in \[9 12 18 22 28 36 45 55 65 82\] C:\\golang>

## 8) Golang *SearchStrings* function \[ascending order\]

The SearchStrings function searches the position of x in a sorted slice of string and returns the index as specified by Search. This function works if slice is in sort order only. If found x in strSlice then it returns index position of strSlice otherwise it returns index position where x fits in sorted slice. The following example shows the usage of SearchStrings() function:

###### Syntax:

func SearchStrings(strSlice \[\]string, x string) int

###### Example:

package main

import (
	"fmt"
	"sort"
)

func main() {
	// string slice in unsorted order
	strSlice := \[\]string{"Texas","Washington","Montana","Alaska","Indiana","Ohio","Nevada"}
	x := "Montana"
	pos := sort.SearchStrings(strSlice,x)
	fmt.Printf("Found %s at index %d in %v\\n", x, pos, strSlice)

	// slice need to be sort in ascending order before to use SearchStrings
	sort.Strings(strSlice)	// slice sorted
	pos = sort.SearchStrings(strSlice,x)
	fmt.Printf("Found %s at index %d in %v\\n", x, pos, strSlice)

	x = "Missouri"
	pos = sort.SearchStrings(strSlice,x)
	fmt.Printf("Found %s at index %d in %v\\n", x, pos, strSlice)

	x = "Utah"
	pos = sort.SearchStrings(strSlice,x)
	fmt.Printf("Found %s at index %d in %v\\n", x, pos, strSlice)

	x = "Ohio"
	pos = sort.SearchStrings(strSlice,x)
	fmt.Printf("Found %s at index %d in %v\\n", x, pos, strSlice)

	x = "OHIO"
	pos = sort.SearchStrings(strSlice,x)
	fmt.Printf("Found %s at index %d in %v\\n", x, pos, strSlice)

	x = "ohio"
	pos = sort.SearchStrings(strSlice,x)
	fmt.Printf("Found %s at index %d in %v\\n", x, pos, strSlice)
}

###### When you run the program, you get the following output:

C:\\golang>go run sort.go Found Montana at index 5 in \[Texas Washington Montana Alaska Indiana Ohio Nevada\] Found Montana at index 2 in \[Alaska Indiana Montana Nevada Ohio Texas Washington\] Found Missouri at index 2 in \[Alaska Indiana Montana Nevada Ohio Texas Washington\] Found Utah at index 6 in \[Alaska Indiana Montana Nevada Ohio Texas Washington\] Found Ohio at index 4 in \[Alaska Indiana Montana Nevada Ohio Texas Washington\] Found OHIO at index 4 in \[Alaska Indiana Montana Nevada Ohio Texas Washington\] Found ohio at index 7 in \[Alaska Indiana Montana Nevada Ohio Texas Washington\] C:\\golang>

## 9) Golang *SearchFloat64s* function \[ascending order\]

The SearchFloat64s function searches the position of x in a sorted slice of float64 and returns the index as specified by Search. This function works if a slice is in sort order only. If found x in fltSlice then it returns index position of fltSlice otherwise it returns index position where x fits in the sorted slice. The following example shows the usage of SearchFloat64s() function:

###### Syntax:

func SearchFloat64s(fltSlice \[\]float64, x float64) int

###### Example:

package main

import (
	"fmt"
	"sort"
)

func main() {
	// string slice in unsorted order
	fltSlice := \[\]float64{962.25, 514.251, 141.214, 96.142, 85.14}
	x := 141.214
	pos := sort.SearchFloat64s(fltSlice,x)
	fmt.Printf("Found %f at index %d in %v\\n", x, pos, fltSlice)

	// slice need to be sort in ascending order before to use SearchFloat64s
	sort.Float64s(fltSlice)	// slice sorted
	pos = sort.SearchFloat64s(fltSlice,x)
	fmt.Printf("Found %f at index %d in %v\\n", x, pos, fltSlice)

	x = 8989.251
	pos = sort.SearchFloat64s(fltSlice,x)
	fmt.Printf("Found %f at index %d in %v\\n", x, pos, fltSlice)

	x = 10.251
	pos = sort.SearchFloat64s(fltSlice,x)
	fmt.Printf("Found %f at index %d in %v\\n", x, pos, fltSlice)

	x = 411.251
	pos = sort.SearchFloat64s(fltSlice,x)
	fmt.Printf("Found %f at index %d in %v\\n", x, pos, fltSlice)

	x = \-411.251
	pos = sort.SearchFloat64s(fltSlice,x)
	fmt.Printf("Found %f at index %d in %v\\n", x, pos, fltSlice)
}

###### When you run the program, you get the following output:

C:\\golang>go run sort.go Found 141.214000 at index 0 in \[962.25 514.251 141.214 96.142 85.14\] Found 141.214000 at index 2 in \[85.14 96.142 141.214 514.251 962.25\] Found 8989.251000 at index 5 in \[85.14 96.142 141.214 514.251 962.25\] Found 10.251000 at index 0 in \[85.14 96.142 141.214 514.251 962.25\] Found 411.251000 at index 3 in \[85.14 96.142 141.214 514.251 962.25\] Found \-411.251000 at index 0 in \[85.14 96.142 141.214 514.251 962.25\] C:\\golang>

## 10) Golang *Search* function \[ascending and descending order\]

The Search function searches the position of x in a sorted slice of string/float/int and returns the index as specified by Search. If found x in data then it returns index position of data otherwise it returns index position where x fits in sorted slice. This function works for both ascending and descending order slice while above 3 search functions only works for ascending order only. The following example shows the usage of Search() function:

###### Syntax:

sort.Search(len(data), func(i int) bool { return data\[i\] >= x })

###### Example:

package main

import (
    "fmt"
    "sort"
)

func main() {

	fmt.Println("\\n######## SearchInts not works in descending order  ######## ")
	intSlice := \[\]int{55, 54, 53, 52, 51, 50, 48, 36, 15, 5}	// sorted slice in descending
    x := 36
    pos := sort.SearchInts(intSlice,x)
    fmt.Printf("Found %d at index %d in %v\\n", x, pos, intSlice)

	fmt.Println("\\n######## Search works in descending order  ########")
	i := sort.Search(len(intSlice), func(i int) bool { return intSlice\[i\] <= x })
	fmt.Printf("Found %d at index %d in %v\\n", x, i, intSlice)

	fmt.Println("\\n\\n######## SearchStrings not works in descending order  ######## ")
	// sorted slice in descending
	strSlice := \[\]string{"Washington","Texas","Ohio","Nevada","Montana","Indiana","Alaska"}
    y := "Montana"
    posstr := sort.SearchStrings(strSlice,y)
    fmt.Printf("Found %s at index %d in %v\\n", y, posstr, strSlice)

	fmt.Println("\\n######## Search works in descending order  ########")
	j := sort.Search(len(strSlice), func(j int) bool {return strSlice\[j\] <= y})
	fmt.Printf("Found %s at index %d in %v\\n", y, j, strSlice)

	fmt.Println("\\n######## Search works in ascending order  ########")
    fltSlice := \[\]float64{10.10, 20.10, 30.15, 40.15, 58.95} // string slice in float64
    z := 40.15
    k := sort.Search(len(fltSlice), func(k int) bool {return fltSlice\[k\] >= z})
	fmt.Printf("Found %f at index %d in %v\\n", z, k, fltSlice)
}

###### When you run the program, you get the following output:

C:\\golang>go run sort.go ######## SearchInts not works in descending order ######## Found 36 at index 0 in \[55 54 53 52 51 50 48 36 15 5\] ######## Search works in descending order ######## Found 36 at index 7 in \[55 54 53 52 51 50 48 36 15 5\] ######## SearchStrings not works in descending order ######## Found Montana at index 0 in \[Washington Texas Ohio Nevada Montana Indiana Alaska\] ######## Search works in descending order ######## Found Montana at index 4 in \[Washington Texas Ohio Nevada Montana Indiana Alaska\] ######## Search works in ascending order ######## Found 40.150000 at index 3 in \[10.1 20.1 30.15 40.15 58.95\] C:\\golang>

## 11) Golang *Sort* function

The Sort function sorts the data interface in both ascending and descending order. It first makes a call to data.Len to determine n, and O(n\*log(n)) calls to data.Less and data.Swap. The following example shows the usage of Sort() function:

###### Syntax:

func Sort(data Interface)

###### Example:

package main

import (
	"fmt"
	"sort"
)

type Mobile struct {
	Brand string
	Price int
}

// ByPrice implements sort.Interface for \[\]Mobile based on
// the Price field.
type ByPrice \[\]Mobile
func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice) Swap(i, j int)      { a\[i\], a\[j\] = a\[j\], a\[i\] }
func (a ByPrice) Less(i, j int) bool { return a\[i\].Price < a\[j\].Price }

// ByBrand implements sort.Interface for \[\]Mobile based on
// the Brand field.
type ByBrand \[\]Mobile
func (a ByBrand) Len() int           { return len(a) }
func (a ByBrand) Swap(i, j int)      { a\[i\], a\[j\] = a\[j\], a\[i\] }
func (a ByBrand) Less(i, j int) bool { return a\[i\].Brand > a\[j\].Brand }

func main() {
	mobile := \[\]Mobile{
		{"Sony", 952},
		{"Nokia", 468},
		{"Apple", 1219},
		{"Samsung", 1045},
	}
	fmt.Println("\\n######## Before Sort #############\\n")
	for \_, v := range mobile {
		fmt.Println(v.Brand, v.Price)
	}

	fmt.Println("\\n\\n######## Sort By Price \[ascending\] ###########\\n")
	sort.Sort(ByPrice(mobile))
	for \_, v := range mobile {
		fmt.Println(v.Brand, v.Price)
	}

	fmt.Println("\\n\\n######## Sort By Brand \[descending\] ###########\\n")
	sort.Sort(ByBrand(mobile))
	for \_, v := range mobile {
		fmt.Println(v.Brand, v.Price)
	}
}

###### When you run the program, you get the following output:

C:\\golang>go run sort.go ######## Before Sort ############# Sony 952 Nokia 468 Apple 1219 Samsung 1045 ######## Sort By Price \[ascending\] ########### Nokia 468 Sony 952 Samsung 1045 Apple 1219 ######## Sort By Brand \[descending\] ########### Sony 952 Samsung 1045 Nokia 468 Apple 1219 C:\\golang>

## 12) Golang *IsSorted* function

The IsSorted function reports whether data is sorted by returns true or false. The following example shows the usage of IsSorted() function:

###### Syntax:

func IsSorted(data Interface) bool

###### Example:

package main

import (
	"fmt"
	"sort"
)

type Mobile struct {
	Brand string
	Price int
}

// ByPrice implements sort.Interface for \[\]Mobile based on
// the Price field.
type ByPrice \[\]Mobile
func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice) Swap(i, j int)      { a\[i\], a\[j\] = a\[j\], a\[i\] }
func (a ByPrice) Less(i, j int) bool { return a\[i\].Price < a\[j\].Price }

func main() {
	mobile1 := \[\]Mobile{
		{"Sony", 952},
		{"Nokia", 468},
		{"Apple", 1219},
		{"Samsung", 1045},
	}
	fmt.Println("\\nFound mobile1 price is sorted :", sort.IsSorted(ByPrice(mobile1)))	// false

	mobile2 := \[\]Mobile{
		{"Sony", 452},
		{"Nokia", 768},
		{"Apple", 919},
		{"Samsung", 1045},
	}
	fmt.Println("\\nFound mobile2 price is sorted :", sort.IsSorted(ByPrice(mobile2)))	// true
}

###### When you run the program, you get the following output:

C:\\golang>go run sort.go Found mobile1 price is sorted : false Found mobile2 price is sorted : true C:\\golang>

## 13) Golang *Slice* function

This Slice function sorts the provided slice given the provided less function. The function panics if the provided interface is not a slice. The following example shows the usage of Slice() function:

###### Syntax:

func Slice(slice interface{}, less func(i, j int) bool)

###### Example:

package main

import (
	"fmt"
	"sort"
)

func main() {
	mobile := \[\]struct {
		Brand string
		Price  int
	}{
		{"Nokia", 700},
		{"Samsung", 505},
		{"Apple", 924},
		{"Sony", 655},
	}
	sort.Slice(mobile, func(i, j int) bool { return mobile\[i\].Brand < mobile\[j\].Brand })
	fmt.Println("\\n\\n######## Sort By Brand \[ascending\] ###########\\n")
    for \_, v := range mobile {
        fmt.Println(v.Brand, v.Price)
    }

	sort.Slice(mobile, func(i, j int) bool { return mobile\[i\].Brand > mobile\[j\].Brand })
	fmt.Println("\\n\\n######## Sort By Brand \[descending\] ###########\\n")
    for \_, v := range mobile {
        fmt.Println(v.Brand, v.Price)
    }

	sort.Slice(mobile, func(i, j int) bool { return mobile\[i\].Price < mobile\[j\].Price })
	fmt.Println("\\n\\n######## Sort By Price \[ascending\] ###########\\n")
    for \_, v := range mobile {
        fmt.Println(v.Brand, v.Price)
    }

	mobile = \[\]struct {
		Brand string
		Price  int
	}{
		{"MI", 900},
		{"OPPO", 305},
		{"iPhone", 924},
		{"sony", 655},
	}

	sort.Slice(mobile, func(i, j int) bool { return mobile\[i\].Brand < mobile\[j\].Brand })
	fmt.Println("\\n\\n######## Sort By Brand \[ascending\] ###########\\n")
    for \_, v := range mobile {
        fmt.Println(v.Brand, v.Price)
    }

}

###### When you run the program, you get the following output:

C:\\golang>go run sort.go ######## Sort By Brand \[ascending\] ########### Apple 924 Nokia 700 Samsung 505 Sony 655 ######## Sort By Brand \[descending\] ########### Sony 655 Samsung 505 Nokia 700 Apple 924 ######## Sort By Price \[ascending\] ########### Samsung 505 Sony 655 Nokia 700 Apple 924 ######## Sort By Brand \[ascending\] ########### MI 900 OPPO 305 iPhone 924 sony 655 C:\\golang>

## 14) Golang *SliceIsSorted* function

This SliceIsSorted function tests whether a slice is sorted. It returns true is data is sorted or false. The following example shows the usage of SliceIsSorted() function:

###### Syntax:

func SliceIsSorted(slice interface{}, less func(i, j int) bool) bool

###### Example:

package main

import (
	"fmt"
	"sort"
)

func main() {
	mobile := \[\]struct {
		Brand string
		Price  int
	}{
		{"Nokia", 700},
		{"Samsung", 505},
		{"Apple", 924},
		{"Sony", 655},
	}
	result := sort.SliceIsSorted(mobile, func(i, j int) bool { return mobile\[i\].Price < mobile\[j\].Price })
	fmt.Println("Found price sorted:", result) // false

	mobile = \[\]struct {
		Brand string
		Price  int
	}{
		{"Nokia", 700},
		{"Samsung", 805},
		{"Apple", 924},
		{"Sony", 955},
	}
	result = sort.SliceIsSorted(mobile, func(i, j int) bool { return mobile\[i\].Price < mobile\[j\].Price })
	fmt.Println("Found price sorted:", result) // true

	mobile = \[\]struct {
		Brand string
		Price  int
	}{
		{"iPhone", 900},
		{"MI", 805},
		{"OPPO", 724},
		{"Sony", 655},
	}
	result = sort.SliceIsSorted(mobile, func(i, j int) bool { return mobile\[i\].Brand < mobile\[j\].Brand })
	fmt.Println("Found brand sorted:", result) // false
}

###### When you run the program, you get the following output:

C:\\golang>go run sort.go Found price sorted: false Found price sorted: true Found brand sorted: false C:\\golang>

## 15) Golang *IntSlice*

IntSlice attaches the methods of Interface to \[\]int, sorting in increasing order. Len used to find the length of slice. Search returns the result of applying SearchInts to the receiver and x. Sort used to sort the slice. The following example shows the usage of IntSlice() function:

###### Syntax:

type IntSlice \[\]int

###### Example:

package main

import (
	"fmt"
	"sort"
)

func main() {
	s := \[\]int{9, 22, 54, 33, \-10, 40} // unsorted
	sort.Sort(sort.IntSlice(s))
	fmt.Println(s)	// sorted
	fmt.Println("Length of Slice: ", sort.IntSlice.Len(s))	// 6
	fmt.Println("40 found in Slice at position: ", sort.IntSlice(s).Search(40))		//	4
	fmt.Println("82 found in Slice at position: ", sort.IntSlice(s).Search(82))		//	6
	fmt.Println("6 found in Slice at position: ", sort.IntSlice(s).Search(6))		//	0
}

###### When you run the program, you get the following output:

C:\\golang>go run sort.go \[\-10 9 22 33 40 54\] Length of Slice: 6 40 found in Slice at position: 4 82 found in Slice at position: 6 6 found in Slice at position: 1 C:\\golang>

## 16) Golang *StringSlice*

StringSlice attaches the methods of Interface to \[\]string, sorting in increasing order. Len used to find the length of slice. Search returns the result of applying SearchStrings to the receiver and x. Sort used to sort the slice. The following example shows the usage of StringSlice function:

###### Syntax:

type StringSlice \[\]string

###### Example:

package main

import (
	"fmt"
	"sort"
)

func main() {
	s := \[\]string{"Washington","Texas","Ohio","Nevada","Montana","Indiana","Alaska"} // unsorted
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)	// sorted
	fmt.Println("Length of Slice: ", sort.StringSlice.Len(s))	// 7
	fmt.Println("Texas found in Slice at position: ", sort.StringSlice(s).Search("Texas"))		//	5
	fmt.Println("Montana found in Slice at position: ", sort.StringSlice(s).Search("Montana"))	//	2
	fmt.Println("Utah found in Slice at position: ", sort.StringSlice(s).Search("Utah"))		//	6

	fmt.Println("OHIO found in Slice at position: ", sort.StringSlice(s).Search("OHIO"))		//	4
	fmt.Println("Ohio found in Slice at position: ", sort.StringSlice(s).Search("Ohio"))		//	4
	fmt.Println("ohio found in Slice at position: ", sort.StringSlice(s).Search("ohio"))		//	7
}

###### When you run the program, you get the following output:

C:\\golang>go run sort.go \[Alaska Indiana Montana Nevada Ohio Texas Washington\] Length of Slice: 7 Texas found in Slice at position: 5 Montana found in Slice at position: 2 Utah found in Slice at position: 6 OHIO found in Slice at position: 4 Ohio found in Slice at position: 4 ohio found in Slice at position: 7 C:\\golang>

## 17) Golang *Float64Slice*

Float64Slice attaches the methods of Interface to \[\]float64, sorting in increasing order. Len used to find the length of slice. Search returns the result of applying SearchFloat64s to the receiver and x. Sort used to sort the slice. The following example shows the usage of Float64Slice function:

###### Syntax:

type Float64Slice \[\]float64

###### Example:

package main

import (
	"fmt"
	"sort"
)

func main() {
	s := \[\]float64{85.201, 14.74, 965.25, 125.32, 63.14} // unsorted
	sort.Sort(sort.Float64Slice(s))
	fmt.Println(s)	// sorted
	fmt.Println("Length of Slice: ", sort.Float64Slice.Len(s))	// 5
	fmt.Println("123.32 found in Slice at position: ", sort.Float64Slice(s).Search(125.32))		//	3
	fmt.Println("999.15 found in Slice at position: ", sort.Float64Slice(s).Search(999.15))		//	5
	fmt.Println("12.14 found in Slice at position: ", sort.Float64Slice(s).Search(12.14))		//	0
}

###### When you run the program, you get the following output:

C:\\golang>go run sort.go \[14.74 63.14 85.201 125.32 965.25\] Length of Slice: 5 123.32 found in Slice at position: 3 999.15 found in Slice at position: 5 12.14 found in Slice at position: 0 C:\\golang>

## 18) Golang *Reverse* function \[descending order\]

The Reverse function returns an slice in the reverse order. The following example shows the usage of Reverse() function:

###### Syntax:

func Reverse(data Interface) Interface

###### Example:

package main

import (
	"fmt"
	"sort"
)

func main() {
	a := \[\]int{15, 4, 33, 52, 551, 90, 8, 16, 15, 105}    // unsorted
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println("\\n",a)

	a = \[\]int{\-15, \-4, \-33, \-52, \-551, \-90, \-8, \-16, \-15, \-105}     // unsorted
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println("\\n",a)

	b := \[\]string{"Montana","Alaska","Indiana","Nevada","Washington","Ohio","Texas"}   // unsorted
	sort.Sort(sort.Reverse(sort.StringSlice(b)))
	fmt.Println("\\n",b)

	b = \[\]string{"ALASKA","indiana","OHIO","Nevada","Washington","TEXAS","Montana"}  // unsorted
	sort.Sort(sort.Reverse(sort.StringSlice(b)))
	fmt.Println("\\n",b)

	c := \[\]float64{90.10, 80.10, 160.15, 40.15, 8.95} //	unsorted
	sort.Sort(sort.Reverse(sort.Float64Slice(c)))
	fmt.Println("\\n",c)

	c = \[\]float64{\-90.10, \-80.10, \-160.15, \-40.15, \-8.95} // unsorted
	sort.Sort(sort.Reverse(sort.Float64Slice(c)))
	fmt.Println("\\n",c)
}

###### When you run the program, you get the following output:

C:\\golang>go run sort.go \[551 105 90 52 33 16 15 15 8 4\] \[\-4 \-8 \-15 \-15 \-16 \-33 \-52 \-90 \-105 \-551\] \[Washington Texas Ohio Nevada Montana Indiana Alaska\] \[indiana Washington TEXAS OHIO Nevada Montana ALASKA\] \[160.15 90.1 80.1 40.15 8.95\] \[\-8.95 \-40.15 \-80.1 \-90.1 \-160.15\] C:\\golang>
