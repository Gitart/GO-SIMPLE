## Как добавить текст в файл на Голанге?
Функция OpenFile с помощью os.O_APPEND открывает уже существующий файл test.txt в режиме добавления. Функция WriteString записывает содержимое в конец файла.
```golang
package main
 
import (
    "io/ioutil"
    "log"
    "fmt"
    "os"
)
 
func AppendFile() {     
    file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        log.Fatalf("failed opening file: %s", err)
    }
    defer file.Close()
 
    len, err := file.WriteString(" The Go language was conceived in September 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google.")
    if err != nil {
        log.Fatalf("failed writing to file: %s", err)
    }
    fmt.Printf("\nLength: %d bytes", len)
    fmt.Printf("\nFile Name: %s", file.Name())
}
 
func ReadFile() {
    data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        log.Panicf("failed reading data from file: %s", err)
    }
    fmt.Printf("\nLength: %d bytes", len(data))
    fmt.Printf("\nData: %s", data)
    fmt.Printf("\nError: %v", err)
}
 
func main() {
    fmt.Printf("######## Append file #########\n")
    AppendFile()
     
    fmt.Printf("\n\n######## Read file #########\n")
    ReadFile()
}
```


## Как читать / писать из / в файл в Голанге?
В нижеприведенной программе функция WriteString используется для записи содержимого в текстовый файл, а функция ReadFile - для чтения содержимого из текстового файла. Программа создаст файл test.txt, если он не существует, или обрежет, если он уже существует.

Golang прочитал файл построчный пример:

```golang
package main
 
import (
    "io/ioutil"
    "log"
    "fmt"
    "os"
)
 
func CreateFile() {
    file, err := os.Create("test.txt") // Truncates if file already exists, be careful!
    if err != nil {
        log.Fatalf("failed creating file: %s", err)
    }
    defer file.Close() // Make sure to close the file when you're done
 
    len, err := file.WriteString("The Go Programming Language, also commonly referred to as Golang, is a general-purpose programming language, developed by a team at Google.")
 
    if err != nil {
        log.Fatalf("failed writing to file: %s", err)
    }
    fmt.Printf("\nLength: %d bytes", len)
    fmt.Printf("\nFile Name: %s", file.Name())
}
 
func ReadFile() {
    data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        log.Panicf("failed reading data from file: %s", err)
    }
    fmt.Printf("\nLength: %d bytes", len(data))
    fmt.Printf("\nData: %s", data)
    fmt.Printf("\nError: %v", err)
}
 
func main() {
    fmt.Printf("########Create a file and Write the content #########\n")
    CreateFile()
     
    fmt.Printf("\n\n########Read file #########\n")
    ReadFile()
}
```
