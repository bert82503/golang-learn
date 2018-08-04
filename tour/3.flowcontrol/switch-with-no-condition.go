package main

import (
  "time"
  "fmt"
)

func main() {
  //t := time.Now()
  //switch {
  //case t.Hour() < 12:
  //  fmt.Println("Good morning!")
  //case t.Hour() < 17:
  //  fmt.Println("Good afternoon.")
  //default:
  //  fmt.Println("Good evening.")
  //}

  h := time.Now().Hour()
  switch {
  case h < 12:
    fmt.Println("Good morning!")
  case h < 17:
    fmt.Println("Good afternoon.")
  default:
    fmt.Println("Good evening.")
  }
}
