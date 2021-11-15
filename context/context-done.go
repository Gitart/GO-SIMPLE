package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {

	// obtain a first background
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	// cancel the context, this could happen in
	// the same goroutine.
	go cancel()

	<-ctx.Done()

	errCancel := ctx.Err()

	if errCancel != nil {
		switch {
		case errors.Is(errCancel, context.Canceled):
			fmt.Println("The context was cancelled")
		case errors.Is(errCancel, context.DeadlineExceeded):
			fmt.Println("The deadline of the context expired")
		}
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Millisecond)

	<-ctx.Done()

	errDeadline := ctx.Err()
	if errDeadline != nil {
		switch {
		case errors.Is(errDeadline, context.Canceled):
			fmt.Println("The context was cancelled")
		case errors.Is(errDeadline, context.DeadlineExceeded):
			fmt.Println("The deadline of the context expired")
		}
	}

}
