package main

import "fmt"

// for is Go’s only looping construct.
func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}
	// 1
	// 2
	// 3

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	// 7
	// 8
	// 9

	for {
		fmt.Println("loop")
		break
	}
	// loop

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	// 1
	// 3
	// 5
}
