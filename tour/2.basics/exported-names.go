package main

import (
	"fmt"
	"math"
)

func main() {
	// # command-line-arguments
	// .\exported-names.go:11:19: undefined: math.pi
	//fmt.Println(math.pi)

	fmt.Println(math.Pi)
}
