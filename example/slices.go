package main

import "fmt"

// Slices are a key data type, giving a more powerful interface to sequences than arrays.
// more details on the design and implementation of slices in Go.
// Go slices usage and internals, https://blog.golang.org/2011/01/go-slices-usage-and-internals.html
func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)
	// emp: [  ]
	// 与数组不同

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	// set: [a b c]
	fmt.Println("get:", s[2])
	// get: c

	fmt.Println("len:", len(s))
	// len: 3

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	// apd: [a b c d e f]
	// 追加操作，动态扩容

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	// cpy: [a b c d e f]

	// 子列表
	l := s[2:5]
	fmt.Println("sl1:", l)
	// sl1: [c d e]

	l = s[:5]
	fmt.Println("sl2:", l)
	// sl2: [a b c d e]

	l = s[2:]
	fmt.Println("sl3:", l)
	// sl3: [c d e f]

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	// dcl: [g h i]

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	// 2d:  [[0] [1 2] [2 3 4]]
	// 二维数组长度允许不相同
}
