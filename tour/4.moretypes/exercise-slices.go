package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
  length := 256
  xy := make([][]uint8, length)
  for i := range xy {
    xy[i] = make([]uint8, length)
    for j := range xy[i] {
      xy[i][j] = rgb(dx, dy)
    }
  }
  return xy
}

func rgb(dx, dy int) uint8 {
  //return uint8((dx + dy) / 2)
  //return uint8(dx * dy)
  return uint8(dx ^ dy)
}

func main() {
  pic.Show(Pic)
}
