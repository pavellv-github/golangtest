package main

import (
	"fmt"
)

func main() {
	myArray := [3]int{1, 2, 3}
	fmt.Println(myArray[2])

	for i := 0; i < len(myArray); i++ {
		item := myArray[i]
		fmt.Println("test a", item)
	}

	for i := 0; i < 10; i++ {
		item := myArray[i]
		fmt.Println("test b", item)
	}

	myArray2 := [...]int{1, 2, 3, 4, 5, 6, 7}

	for i := 0; i < 10; i++ {
		item := myArray2[i]
		fmt.Println("test c", item)
	}
}
