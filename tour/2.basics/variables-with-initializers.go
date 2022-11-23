package main

import "fmt"

// var i, j int = 1, 2
var i, j = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	// 1 2 true false no!
	fmt.Println(i, j, c, python, java)
}
