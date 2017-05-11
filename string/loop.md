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


