// 程序启动入口
package main

import (
	"fmt"
	"rsc.io/quote"
)

// Getting started - Say Hello, World with Go.
// Tutorial: Get started with Go
// https://golang.google.cn/doc/tutorial/getting-started

func main() {
	// Write some code
	fmt.Println("Hello, World!")
	fmt.Println("你好，世界！")

	// Call code in an external package
	fmt.Println(quote.Go())
}
