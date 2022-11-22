
# Generate file

```golang
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func validLines(filename string) (func() (string, bool)) {

    file, _ := os.Open(filename)
    scanner := bufio.NewScanner(file)

    return func() (string, bool) {
        buff := ""
        for scanner.Scan() {
            line := scanner.Text()
            line = strings.TrimSpace(line)

            if line == "" {
                continue
            }
            buff += line
            if line[len(line)-1] != ';' {
                continue
            }
            return buff, true
        }

        file.Close()
        return "", false
    }
}

func main() {
    vline := validLines("myfile.txt")
    for line, ok := vline(); ok; {
        fmt.Println(line)
    }
}
```

## Output

```
#123= FOOBAR(1.,'text');
#123= FOOBAR(1.,'text');
#123= FOOBAR(1.,'text');
#123= FOOBAR(1.,'text');
#123= FOOBAR(1.,'text');
#123= FOOBAR(1.,'text');
#123= FOOBAR(1.,'text');
...
```

