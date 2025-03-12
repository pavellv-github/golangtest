package main

import (
	"fmt"
)

func main() {
	//  myMap := make(map[string]int);

	myMap := map[string]int{
		"item1": 1,
		"item2": 2,
		"item3": 3,
	}

	fmt.Println(myMap["item1"])

	delete(myMap, "item2")

	fmt.Println(myMap["item2"])

	myMap["item2"] = 222

	fmt.Println(myMap["item2"])

	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
