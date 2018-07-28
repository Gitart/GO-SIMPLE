## Alternate reflex

```go
package main

import (
    "github.com/davecgh/go-spew/spew"
)

type Project struct {
    Id      int64  `json:"project_id"`
    Title   string `json:"title"`
    Name    string `json:"name"`
    Data    string `json:"data"`
    Commits string `json:"commits"`
}

func main() {

    o := Project{Name: "hello", Title: "world"}
    spew.Dump(o)
}
```

### output:

```
(main.Project) {
 Id: (int64) 0,
 Title: (string) (len=5) "world",
 Name: (string) (len=5) "hello",
 Data: (string) "",
 Commits: (string) ""
}
```
