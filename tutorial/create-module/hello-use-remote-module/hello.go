// 应用程序启动入口
package main

import (
	"fmt"
	"github.com/bert82503/golang-learn/tutorial/create-module/greetings"
)

// Call your code from another module
// https://golang.google.cn/doc/tutorial/call-module-code

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
