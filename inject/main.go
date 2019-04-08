package main

import (
	"fmt"
)

// BuildTime inject flag when go build
var BuildTime string

func main() {
	fmt.Printf("Hello world, version: %+v\n", BuildTime)
}
