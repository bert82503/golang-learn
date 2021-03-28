package main

import (
	"fmt"
	"math"
)

const s string = "constant"

// Go supports constants of character, string, boolean, and numeric values.
func main() {
	fmt.Println(s)
	// constant

	const n = 5_0000_0000

	const d = 3e20 / n
	fmt.Println(d)
	// 6e+11

	fmt.Println(int64(d))
	// 600000000000

	fmt.Println(math.Sin(n))
	// -0.28470407323754404
}
