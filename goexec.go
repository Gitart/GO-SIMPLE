/* ... <== see fragment description ... */

package main

import (
    "fmt"
    "log"
    "os/exec"
)

func main() {
    out, err := exec.Command("ls").Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("The ls result is:\n%s", out)
}

/* Expected Output:
The ls result is:
output.exe
output.go
*/
