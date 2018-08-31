// Context is used to cancel(stop) gorutines without leaking (continueing gorutines in background) when the process is canceled

package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Println("------------------")
	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	/*
		ctx, _ = context.WithCancel(ctx)
		fmt.Println("------------------")
		fmt.Println("context:\t", ctx)
		fmt.Println("context err:\t", ctx.Err())
		fmt.Printf("context type:\t%T\n", ctx)
	*/
	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("------------------")
	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("context err:\t", cancel)
	fmt.Printf("context type:\t%T\n", cancel)
	cancel()

	fmt.Println("------------------")
	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("context err:\t", cancel)
	fmt.Printf("context type:\t%T\n", cancel)

}


