

# Read file

```golang
package myPackage

import (
    "io/ioutil"
    "strings"
)

func FileContainsName(filename string, name string) (bool err) {
    file, _ := ioutil.ReadFile(filename)
    content := string(file)

    return strings.Contains(content, name), err
}
```
