package main

import (
	"fmt"
	"math"
)

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	// 这里开始就不能使用 v 了
	return lim
}

/*
27 >= 20
9 20
*/
func main() {
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}
