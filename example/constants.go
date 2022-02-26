package main

import (
	"fmt"
	"math"
)

// 声明一个常量值
const s string = "constant"

/*
Go supports constants of character, string, boolean, and numeric values.

const declares a constant value.

A const statement can appear anywhere a var statement can.

Constant expressions perform arithmetic with arbitrary precision.

A numeric constant has no type until it’s given one, such as by an explicit conversion.

A number can be given a type by using it in a context that requires one, such as a variable assignment or function call.
For example, here math.Sin expects a float64.
*/
func main() {
	fmt.Println(s)
	// constant

	const n = 500_000_000

	const d = 3e20 / n
	fmt.Println(d)
	// 6e+11

	fmt.Println(int64(d))
	// 600000000000

	fmt.Println(math.Sin(n))
	// -0.28470407323754404
}
