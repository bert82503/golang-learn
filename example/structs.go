package main

import "fmt"

// 结构体是字段的类型化集合
// 对象
type person struct {
	// fields，字段，状态
	name string
	age  int
}

// It’s idiomatic to encapsulate new struct creation in constructor functions
func newPerson(name string) *person {
	// You can safely return a pointer to local variable as a local variable will survive the scope of the function.
	p := person{name: name}
	p.age = 42
	return &p
}

// Go’s structs are typed collections of fields.
// They’re useful for grouping data together to form records.
// Go的结构体是字段的类型化集合。它们对于将数据分组形成记录非常有用。
func main() {
	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})
	// {Bob 20}

	fmt.Println(person{name: "Alice", age: 30})
	// {Alice 30}

	fmt.Println(person{name: "Fred"})
	// {Fred 0}

	// An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})
	// &{Ann 40}

	fmt.Println(newPerson("Jon"))
	// &{Jon 42}

	// Access struct fields with a dot.
	s := person{name: "sean", age: 50}
	fmt.Println(s.name)
	// sean

	// You can also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age)
	// 50

	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age)
	// 51
	fmt.Println(s.age)
	// 51
}
