package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	switch result := 8; {
	case result == 12:
		fmt.Println("a")
	case result == 10-2:
		fmt.Println("b")
	case result > 100:
		fmt.Println("c")
	case result%2 == 12:
		fmt.Println("e")
	default:
		fmt.Println(result)
	}
}
