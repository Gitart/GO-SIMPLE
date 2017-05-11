## Прокрутка линий из файла


```golang 

// Loop file by lines and words
// http://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go

package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func main() {
    file, err := os.Open("./tmp/dat")
    if err != nil {
       log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    i:=0
    s:=""





    for scanner.Scan() {
        i++
        s=scanner.Text()
        
        // сканирование линии по словам
        scanerword(s) 
        
        // Catching lines with prefix rom
        if strings.Contains(s, "rom"){
           fmt.Println("========",i,s)
        }
        

        if strings.Contains(s, "rm"){
           fmt.Println("++++++++++++++",i,s)
        }



    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}





func createFileWithLongLine(fn string) (err error) {
    file, err := os.Create(fn)
    defer file.Close()

    if err != nil {
        return err
    }

    w := bufio.NewWriter(file)

    fs := 1024 * 1024 * 4 // 4MB

    // Create a 4MB long line consisting of the letter a.
    for i := 0; i < fs; i++ {
        w.WriteRune('a')
    }

    // Terminate the line with a break.
    w.WriteRune('\n')

    // Put in a second line, which doesn't have a linebreak.
    w.WriteString("Second line.")

    w.Flush()

    return
}

func limitLength(s string, length int) string {
    if len(s) < length {
        return s
    }

    return s[:length]
}




func scanerword(words string){
      // An artificial input source.
    // const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"

    scanner := bufio.NewScanner(strings.NewReader(words))
    // Set the split function for the scanning operation.
    scanner.Split(bufio.ScanWords)
    // Count the words.
    count := 0
    for scanner.Scan() {
        fmt.Println(scanner.Text())
        count++
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading input:", err)
    }
    fmt.Printf("%d\n", count)
}
```


```golang
func readLine(path string) {
  inFile, _ := os.Open(path)
  defer inFile.Close()
  scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines) 
  
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
}
```


## Sample
```golang
// Reading and writing files are basic tasks needed for
// many Go programs. First we'll look at some examples of
// reading files.

package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // Perhaps the most basic file reading task is
    // slurping a file's entire contents into memory.
    dat, err := ioutil.ReadFile("/tmp/dat")
    check(err)
    fmt.Print(string(dat))

    // You'll often want more control over how and what
    // parts of a file are read. For these tasks, start
    // by `Open`ing a file to obtain an `os.File` value.
    f, err := os.Open("/tmp/dat")
    check(err)

    // Read some bytes from the beginning of the file.
    // Allow up to 5 to be read but also note how many
    // actually were read.
    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))

    // You can also `Seek` to a known location in the file
    // and `Read` from there.
    o2, err := f.Seek(6, 0)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

    // The `io` package provides some functions that may
    // be helpful for file reading. For example, reads
    // like the ones above can be more robustly
    // implemented with `ReadAtLeast`.
    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    // There is no built-in rewind, but `Seek(0, 0)`
    // accomplishes this.
    _, err = f.Seek(0, 0)
    check(err)

    // The `bufio` package implements a buffered
    // reader that may be useful both for its efficiency
    // with many small reads and because of the additional
    // reading methods it provides.
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))

    // Close the file when you're done (usually this would
    // be scheduled immediately after `Open`ing with
    // `defer`).
    f.Close()

}
```



