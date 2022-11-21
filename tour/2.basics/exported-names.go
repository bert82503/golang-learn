package main

import (
	"fmt"
	"math"
)

func main() {
	// # command-line-arguments
	// .\exported-names.go:11:19: undefined: math.pi
	//fmt.Println(math.pi)

	// 3.141592653589793
	fmt.Println(math.Pi)
}
