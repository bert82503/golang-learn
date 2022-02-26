package main

import "fmt"

/*
Go has various value types including strings, integers, floats, booleans, etc.
Here are a few basic examples.

Strings, which can be added together with +.

Integers and floats.

Booleans, with boolean operators as you’d expect.
*/
func main() {
	fmt.Println("go" + "lang")
	// golang

	fmt.Println("1+1 =", 1+1)
	// 1+1 = 2
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	// 7.0/3.0 = 2.3333333333333335

	fmt.Println(true && false)
	// false
	fmt.Println(true || false)
	// true
	fmt.Println(!true)
	// false
}
