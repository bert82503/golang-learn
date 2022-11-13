// 应用程序启动入口
package main

import (
	"fmt"
	"github.com/bert82503/golang-learn/tutorial/create-module/greetings"
	"log"
)

// 2. Call your code from another module
// https://golang.google.cn/doc/tutorial/call-module-code

//func main() {
//	// Get a greeting message and print it.
//	message := greetings.Hello("Gladys")
//	fmt.Println(message)
//}

// 3. Return and handle an error
// https://golang.google.cn/doc/tutorial/handle-errors
// 3.2. In your hello/hello.go file, handle the error now returned by the Hello function,
// along with the non-error value.

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	// 3. Return and handle an error
	//message, err := greetings.Hello("")
	// 4. Return a random greeting
	// https://golang.google.cn/doc/tutorial/random-greeting
	// You're just adding Gladys's name (or a different name, if you like)
	// as an argument to the Hello function call in hello.go.
	message, err := greetings.Hello("Gladys")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
