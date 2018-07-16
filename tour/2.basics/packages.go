package main

import (
  "fmt"
  "math/rand"
)

func main()  {
  rand.Seed(23)
  fmt.Println("我最喜欢的数字是", rand.Intn(100))
}
