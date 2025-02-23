package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello world")
}

func greet(name string) {
	fmt.Println("Hello", name)
}

func add(a, b int) int {
	return a + b
}

func main() {
	sayHello()
	greet("Pavel")
	result := add(5, 10)
	fmt.Println(result)
}
