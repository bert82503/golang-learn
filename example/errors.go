package main

import (
	"errors"
	"fmt"
)

// By convention, errors are the last return value and have type error, a built-in interface.
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a basic error value with the given error message.
		return -1, errors.New("can't work with 42")
	}
	// A nil value in the error position indicates that there was no error.
	return arg + 3, nil
}

// use custom types as errors by implementing the Error() method on them.
type argError struct {
	arg  int
	prob string
}

// 实现error接口
// type error interface {
// Error() string
// }
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// use &argError syntax to build a new struct, supplying values for the two fields arg and prob.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

// In Go it’s idiomatic to communicate errors via an explicit, separate return value.
// Go’s approach makes it easy to see which functions return errors and to handle them using the same language constructs employed for any other, non-error tasks.
func main() {
	// The two loops below test out each of our error-returning functions.
	// Note that the use of an inline error check on the if line is a common idiom in Go code.
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	// f1 worked: 10
	// f1 failed: can't work with 42

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	// f2 worked: 10
	// f2 failed: 42 - can't work with it

	// If you want to programmatically use the data in a custom error,
	// you’ll need to get the error as an instance of the custom error type via type assertion.
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		// 42
		fmt.Println(ae.prob)
		// can't work with it
	}
}
