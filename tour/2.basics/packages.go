package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// To see a different number, seed the number generator; see `rand.Seed`.
	rand.Seed(time.Now().UnixNano())

	// My favorite number is 3
	fmt.Println("My favorite number is", rand.Intn(10))
}
