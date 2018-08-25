package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("runtime: %s\narchitecture: %s", runtime.GOOS, runtime.GOARCH)
	// CLI: GOOS=darwin GOARCH=amd64 go build main.go
}
