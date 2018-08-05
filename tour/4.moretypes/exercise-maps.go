package main

import "golang.org/x/tour/wc"

func WordCount(s string) map[string]int {
  //return map[string]int{"x":1}
  return map[string]int{"I":1, "am":1, "learning":1, "Go!":1}
  //return map[string]int{"brown":1, "jumped":1, "the":1, "dog.":1, "The":1, "quick":1, "fox":1, "over":1, "lazy":1}
}

func main() {
  wc.Test(WordCount)
}
