package main

import "fmt"

// This fact function calls itself until it reaches the base case of fact(0).
// 递归函数
func fact(n int) int {
	// 退出条件
	if n == 0 {
		return 1
	}
	// 递归调用，分而治之思想
	return n * fact(n-1)
}

// Go supports recursive functions.
func main() {
	fmt.Println(fact(7))
	// 5040
}
