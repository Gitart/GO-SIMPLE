https://go.dev/play/p/KSqJin3l0QX
## Sample work Context easy

```go
package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "1", "one")
	ctx = context.WithValue(ctx, "2", "two")
	ctx = context.WithValue(ctx, "3", "3two")

	fmt.Println(ctx.Value("1"))
	fmt.Println(ctx.Value("2"))
	fmt.Println(ctx.Value("3"))

	tr(ctx)

}

func tr(ctx context.Context) {
	ctx = context.WithValue(ctx, "4", "5555555")
	fmt.Println(ctx.Value("4"))
	ters(ctx)
}

func ters(ctx context.Context) {
	ctx = context.WithValue(ctx, "Norm", "Норма выполнения")
	fmt.Println(ctx.Value("Norm"))
}

````
