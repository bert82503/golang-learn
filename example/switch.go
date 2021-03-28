package main

import (
	"fmt"
	"time"
)

// Switch statements express conditionals across many branches.
func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	// basic switch
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	// Write 2 as two

	// use commas to separate multiple expressions in the same case statement
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	// It's a weekday
	// 非周末运行

	// express if/else logic
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	// It's after noon

	// A type switch compares types instead of values.
	// use this to discover the type of an interface value.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	// I'm a bool
	whatAmI(1)
	// I'm an int
	whatAmI("hey")
	// Don't know type string
}
