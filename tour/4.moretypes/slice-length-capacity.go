package main

import "fmt"

func printSlice(s []int) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
  s := []int{2, 3, 5, 7, 11, 13}
  printSlice(s)

  // Slice the slice to give it zero length.
  // len=0 cap=6 []
  s = s[:0]
  printSlice(s)

  // Extend its length.
  // len=4 cap=6 [2 3 5 7]
  s = s[:4]
  printSlice(s)

  // Drop its first two values.
  // len=2 cap=4 [5 7]
  s = s[2:]
  printSlice(s)

  // panic: runtime error: slice bounds out of range
  // len=2 cap=4 [5 7]
  //s = s[0:5]
  //printSlice(s)

  // len=4 cap=4 [5 7 11 13]
  s = s[0:4]
  printSlice(s)

  // len=3 cap=4 [5 7 11]
  s = s[:3]
  printSlice(s)
}
