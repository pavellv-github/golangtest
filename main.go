package main

import (
	"fmt"
	"greetings/greetings"
)

func main() {
	var text string = "Hello world"
	fmt.Println(text)

	message := greetings.Hello("Pavel")
	fmt.Println(message)
}
