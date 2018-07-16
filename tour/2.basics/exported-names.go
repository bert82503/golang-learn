package main

import (
  "fmt"
  "math"
)

func main() {
  // cannot refer to unexported name math.pi
  // undefined: math.pi
  //fmt.Println(math.pi)

  fmt.Println(math.Pi)
}
