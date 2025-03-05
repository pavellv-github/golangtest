package main

import (
	"fmt"
)

var slice = make([]int, 20)

func iterate(slice []int) {
	for i := 0; i < len(slice); i++ {
		item := slice[i]
		fmt.Println("asd", item)
	}
}

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7}

	slice1 := mySlice[:]
	slice2 := mySlice[1:3]
	slice3 := mySlice[4:6]
	slice4 := mySlice[:6]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	mySlice = append(mySlice, 8, 9, 10, 11)
	fmt.Println(mySlice)

	mySlice2 := []int{12, 13, 14, 15}
	mySlice3 := append(mySlice, mySlice2...)
	fmt.Println(mySlice3)

	small := []int{1}
	big := []int{1, 2, 3, 4, 5, 6, 7, 8}
	iterate(small)
	iterate(big)

	board := [][]string{
		[]string{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "0"
	board[1][2] = "X"
	board[1][0] = "0"
	board[0][2] = "X"

	fmt.Println(board)
}
