package main

import "fmt"

// an array is a numbered sequence of elements of a specific length.
func main() {
	var a [5]int
	fmt.Println("emp:", a)
	// emp: [0 0 0 0 0]

	a[4] = 100
	fmt.Println("set:", a)
	// set: [0 0 0 0 100]
	fmt.Println("get:", a[4])
	// get: 100

	fmt.Println("len:", len(a))
	// len: 5

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	// dcl: [1 2 3 4 5]

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	// 2d:  [[0 1 2] [1 2 3]]
}
