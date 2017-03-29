

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


```golang
package gobotto

import (
    "strings"
)

// This is our new interface!
type MyFileReader interface {
    Read(name string) ([]byte, error)
}

// Notice we're not using ioutil's ReadFile anymore, now we can provide our own file reader and implement a fake Read method in it
func FileContainsName(reader MyFileReader, filename string, name string) bool {
    file, _ := reader.Read(filename)
    content := string(file)

    return strings.Contains(content, name)
}
```

Thanks to this interface we will be able to implement a fake version of MyFileReader and pass it to our FileContainsName function. So, this is what our test would look like:


```golang
import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type fakeReader struct{}

// Notice how we're now able to implement any Read method we want to and return arbitrary content
func (fakeReader *fakeReader) Read(filename string) ([]byte, error) {
    return []byte("My friend John is nice."), nil
}

func TestFileContainsName(t *testing.T) {
    assert.True(t, FileContainsName(new(fakeReader), "whatever.txt", "John"))
}
```
