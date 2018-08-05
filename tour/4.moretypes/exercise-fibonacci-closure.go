package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
//斐波那契数，亦称之为斐波那契数列（意大利语： Successione di Fibonacci)，
//又称黄金分割数列、费波那西数列、费波拿契数、费氏数列，指的是这样一个数列：
//0、1、1、2、3、5、8、13、21、……在数学上，斐波纳契数列以如下被以递归的方法定义：
//F0=0，F1=1，Fn=Fn-1+Fn-2（n>=2，n∈N*），用文字来说，就是斐波那契数列列由
//0 和 1 开始，之后的斐波那契数列系数就由之前的两数相加。
func fibonacci() func() int {
  // 预先定义好前两个值
  f0, f1 := 0, 1

  return func() int {
    // 记录f0的值
    fn := f0
    // 重新赋值(这个就是核心代码)
    f0, f1 = f1, (f1 + f0)
    return fn
  }
}

//func fibonacci(n int) int {
//  if n < 2 {
//    return n
//  }
//  return fibonacci(n-2) + fibonacci(n-1)
//}

func main() {
  //for i := 0; i < 10; i++ {
  //  fmt.Printf("%d, ", fibonacci(i))
  //}

  // 返回一个闭包函数
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}
