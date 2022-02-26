package main

import "fmt"

/**
 * Our first program will print the classic “hello world” message.
 * <p>
 * To run the program, put the code in hello-world.go and use go run.
 * $ go run hello-world.go
 * hello world
 * <p>
 * Sometimes we’ll want to build our programs into binaries. We can do this using go build.
 * $ go build hello-world.go
 * $ ls
 * hello-world    hello-world.go
 * <p>
 * We can then execute the built binary directly.
 * $ ./hello-world
 * hello world
 * <p>
 * Now that we can run and build basic Go programs, let’s learn more about the language.
 */
func main() {
	fmt.Println("hello world")
}
