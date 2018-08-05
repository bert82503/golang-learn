package main

import "fmt"

type Vertex2 struct {
  X int
  Y int
}

func main() {
  v := Vertex2{1, 2}
  v.X = 4
  fmt.Println(v.X)
}
