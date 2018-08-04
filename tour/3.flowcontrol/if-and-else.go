package main

import (
  "math"
  "fmt"
)

func pow2(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  } else {
    fmt.Printf("%g >= %g\n", v, lim)
  }
  // 这里开始就不能使用 v 了
  return lim
}

func main() {
  fmt.Println(
    pow2(3, 2, 10),
    pow2(3, 3, 20),
  )
}
