package main

import "fmt"

/*
Branching with if and else in Go is straight-forward.

Here’s a basic example.

You can have an if statement without an else.

A statement can precede conditionals; any variables declared in this statement are available in all branches.

Note that you don’t need parentheses around conditions in Go, but that the braces are required.

There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.
*/
func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	// 7 is odd

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	// 8 is divisible by 4

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	// 9 has 1 digit
}
