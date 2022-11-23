package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// 代码可读性更好
	//v := math.Pow(x, n)
	//if v < lim {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
	// 变量作用域
	// prog.go:12:9: undefined: v
	//return v
}

func main() {
	// 9 20
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
