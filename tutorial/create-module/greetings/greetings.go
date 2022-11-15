package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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
// 3.1. In greetings/greetings.go, add the code highlighted below.

// Hello returns a greeting for the named person.
//func Hello(name string) (string, error) {
//	// If no name was given, return an error with a message.
//	if name == "" {
//		return "", errors.New("empty name")
//	}
//
//	// If a name was received, return a value that embeds the name
//	// in a greeting message.
//	message := fmt.Sprintf("Hi, %v. Welcome!", name)
//	return message, nil
//}

// 4. Return a random greeting
// https://golang.google.cn/doc/tutorial/random-greeting
// You'll add a small slice to contain three greeting messages,
// then have your code return one of the messages randomly.

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// 5. Return greetings for multiple people
// https://golang.google.cn/doc/tutorial/greetings-multiple-people

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in the function.
// Go executes init functions automatically at program startup,
// after global variables have been initialized.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
