package greetings

import "fmt"

// Create a module
// Tutorial: Create a Go module
// https://golang.google.cn/doc/tutorial/create-module
// Start a module that others can use

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	//var message string
	//message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
