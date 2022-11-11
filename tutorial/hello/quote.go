package main

import "fmt"

import "rsc.io/quote"

func main() {
	// Call code in an external package
	fmt.Println(quote.Go())
}
