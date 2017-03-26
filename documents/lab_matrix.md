## Пример лаб по матрице

```golang
package main

import (
       "fmt"
       "math"
       "encoding/json"
)


type Mst map[string]interface{}  
type Msl []interface{}  
type Mil []int64

type Msi struct {
     St string
     Zn []int64 	
}


// 
func main(){
    
   
    var m []Mst    // входная матрица
    var v []Mst    // выходная матрица
    var t []Mst    // выходная матрица

    // Опредление массива матрицы 3*3
    Dt:=[]byte(`[{"A":[5,2,3]}, {"A":[17,2,5]}, {"A":[1,3,8]}]`)     // Исходная матрица 
    Vt:=[]byte(`[{"A":[0,0,0]}, {"A":[0,0,0]},  {"A":[0,0,0]}]`)     // Промежуточные вычисления
    Tr:=[]byte(`[{"A":[0,0,0]}, {"A":[0,0,0]},  {"A":[0,0,0]}]`)     // Трансопнрирование 

	json.Unmarshal(Dt, &m)
    json.Unmarshal(Vt, &v)
    json.Unmarshal(Tr, &t)


    // -110
    F:=GetEl(0, 0, m) * (GetEl(1, 1, m) * GetEl(2, 2, m) - GetEl(1, 2, m) *GetEl(2, 1, m)) - GetEl(0, 1, m) * (GetEl(1, 0, m) * GetEl(2, 2, m)-GetEl(1, 2, m)*GetEl(2, 0, m)) + GetEl(0, 2, m)  *(GetEl(1, 0, m)*GetEl(2, 1, m)-GetEl(1, 1, m)* GetEl(2, 0, m))
    var P float64 =-1


    fmt.Println("Определитель матрицы : ", F)


    // первая строка
    A:= math.Pow(P,(1+1)) * (GetEl(1, 1, m)* GetEl(2, 2, m) - GetEl(1, 2, m) * GetEl(2, 1, m))
    SetEl(0, 0, A, v) 
    fmt.Println("A11 : ",A)

    
    A = math.Pow(P,(1+2)) * (GetEl(1, 0, m)* GetEl(2, 2, m) - GetEl(1, 2, m) * GetEl(2, 0, m))
    SetEl(0, 1, A, v) 
    fmt.Println("A12 : ",A)
   
    A = math.Pow(P,(1+3)) * (GetEl(1, 0, m) * GetEl(2, 1, m) - GetEl(1, 1, m) * GetEl(2, 0, m))
    SetEl(0, 2, A, v) 
    fmt.Println("A13 : ",A)

    // Вторя строка
    A = math.Pow(P,(2+1)) * (GetEl(0, 1, m) * GetEl(2, 2, m) - GetEl(0, 2, m) * GetEl(2, 1, m))
    SetEl(1, 0, A, v) 
    fmt.Println("A21 : ",A)

    A = math.Pow(P,(2+2)) * (GetEl(0, 0, m) * GetEl(2, 2, m) - GetEl(0, 2, m) * GetEl(2, 0, m))
    SetEl(1, 1, A, v)
    fmt.Println("A22 : ",A) 
  
    A = math.Pow(P,(2+3)) * (GetEl(0, 0, m) * GetEl(2, 1, m) - GetEl(0, 1, m) * GetEl(2, 0, m))
    SetEl(1, 2, A, v) 
    fmt.Println("A23 : ",A)
    
    
    // Третья строка
    A = math.Pow(P,(3+1)) * (GetEl(0, 1, m) * GetEl(1, 2, m) - GetEl(0, 2, m) * GetEl(1, 1, m))
    SetEl(2, 0, A, v) 
    fmt.Println("A31 : ",A)
    
    A = math.Pow(P,(3+2)) * (GetEl(0, 0, m) * GetEl(1, 2, m) - GetEl(0, 2, m) * GetEl(1, 0, m))
    SetEl(2, 1, A, v) 
    fmt.Println("A32 : ",A)
    
    A = math.Pow(P,(3+3)) * (GetEl(0, 0, m) * GetEl(1, 1, m) - GetEl(0, 1, m) * GetEl(1, 0, m))
    SetEl(2, 2, A, v) 
    fmt.Println("A33 : ",A)


    fmt.Println("Матрица из алгебраических дополнений элементов : ", A)
    fmt.Println(v)

    // Transponirovanie
    fmt.Println("Транспонированная матрица")
    fmt.Println("")

    // Транспонирование
    for i:=0; i<3;i++ {
    	ti:=1/F
    	// A=GetEl(i, 0, m )
        SetEl(0, i, GetEl(i, 0, v ) * ti, t)
        SetEl(1, i, GetEl(i, 1, v ) * ti, t)
        SetEl(2, i, GetEl(i, 2, v ) * ti, t)

       fmt.Println(GetEl(0, i, v ),GetEl(1, i, v ),GetEl(2, i, v ))  
    }    
    
    fmt.Println("Обратная матрица")
    // Цикл по Элементу А
    for i:=0; i<3;i++ {
        fmt.Println(GetEl(0, i, t ),GetEl(1, i, t ),GetEl(2, i, t ))  
    }    

    // fmt.Println("--------",t, "\n")
    // // SetEl(0, 0, 2222, v) 
    // ff:= GetEl(0, 0, v)
    // // fmt.Println("Промежуточный результат : ", A)   
    //  fmt.Println("Результат : ", F, ff)
}


// Получение элемента по строке элемеенту и положению в массиве
// Str  = положение строки в матрице
// Elm  = положение элемента в строке 
func GetEl(Str, Elm int, Dat []Mst) float64{
     return Dat[Str]["A"].([]interface{})[Elm].(float64) 
}



// Запись в матрицу
func SetEl(Str, Elm int, Res float64, Dat []Mst) {
     Dat[Str]["A"].([]interface{})[Elm]=Res 
}
```


## Пример сложения матрицы
```golang
/*
  {
    "date_of_creation" => "19 Dec 2016, Mon",
    "aim_of_program"   => "Matrix multiplication in Golang",
    "coded_by"         => "Rishikesh Agrawani",
    "Go_version"       => "1.7",
  }
*/
package main
 
import "fmt"
 
func main() {
    //Defining 2D matrices
    m1 := [3][3]int{
        [3]int{12, 6, 1},
        [3]int{13, 7, 1},
        [3]int{1, 5, 1},
    }
    m2 := [3][3]int{
        [3]int{1, 1, 1},
        [3]int{1, 1, 1},
        [3]int{1, 1, 1},
    }
 
    //Declaring a matrix variable for holding the multiplication results
    var m3 [3][3]int
 
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            m3[i][j] = 0
            for k := 0; k < 3; k++ {
                m3[i][j] = m3[i][j] + (m1[i][k] * m2[k][j])
            }
        }
    }
 
    twoDimensionalMatrices := [3][3][3]int{m1, m2, m3}
 
    matrixNames := []string{"MATRIX1", "MATRIX2", "MATRIX3 = MATRIX1*MATRIX2"}
    for index, m := range twoDimensionalMatrices {
        fmt.Println(matrixNames[index], index+1, ":")
        showMatrixElements(m)
        fmt.Println()
    }
}
 
//A function that displays matix elements
func showMatrixElements(m [3][3]int) {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Printf("%d\t", m[i][j])
        }
        fmt.Println()
    }
}
 
 ```
