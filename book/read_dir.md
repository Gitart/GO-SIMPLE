###  ReadAll, ReadDir и ReadFile из пакета ввода-вывода

ReadAll читает от r до ошибки или EOF и возвращает прочитанные данные. Успешный вызов возвращает err == nil, а не err == EOF.
ReadDir читает каталог с именем dirname и возвращает список записей каталога, отсортированных по имени файла.
ReadFile читает файл с именем filename и возвращает его содержимое. Успешный вызов возвращает err == nil, а не err == EOF.
***** Создайте текстовый файл input.txt и напишите содержимое, которое вы хотите прочитать, в том же каталоге, в котором вы будете запускать программу ниже *****

```golang
package main
 
import (
    "io/ioutil"
    "log"
    "fmt"
    "os"
)
 
func exampleReadAll() {
    file, err := os.Open("input.txt")   
    if err != nil {
        log.Panicf("failed reading file: %s", err)
    }
    defer file.Close()
    data, err := ioutil.ReadAll(file)
    fmt.Printf("\nLength: %d bytes", len(data))
    fmt.Printf("\nData: %s", data)
    fmt.Printf("\nError: %v", err)
}
 
func exampleReadDir() {
    entries, err := ioutil.ReadDir(".")
    if err != nil {
        log.Panicf("failed reading directory: %s", err)
    }   
    fmt.Printf("\nNumber of files in current directory: %d", len(entries))  
    fmt.Printf("\nError: %v", err)
}
 
func exampleReadFile() {
    data, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Panicf("failed reading data from file: %s", err)
    }
    fmt.Printf("\nLength: %d bytes", len(data))
    fmt.Printf("\nData: %s", data)
    fmt.Printf("\nError: %v", err)
}
 
func main() {
    fmt.Printf("########Demo of ReadAll function#########\n")
    exampleReadAll()
     
    fmt.Printf("\n\n########Demo of ReadDir function#########\n")
    exampleReadDir()
     
    fmt.Printf("\n\n########Demo of ReadFile function#########\n")
    exampleReadFile()
}
```
