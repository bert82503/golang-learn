package greetings

import (
	"errors"
	"fmt"
)

// 1. Create a module
// Tutorial: Create a Go module
// https://golang.google.cn/doc/tutorial/create-module

// Start a module that others can use

// Hello returns a greeting for the named person.
//func Hello(name string) string {
//	// Return a greeting that embeds the name in a message.
//	message := fmt.Sprintf("Hi, %v. Welcome!", name)
//	// Taking the long way, you might have written this as:
//	//var message string
//	//message = fmt.Sprintf("Hi, %v. Welcome!", name)
//	return message
//}

// 3. Return and handle an error
// https://golang.google.cn/doc/tutorial/handle-errors
// 1. In greetings/greetings.go, add the code highlighted below.

func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
