package main

import (
	"fmt"
)

func main() {
	// var text string = "Hello world"
	// fmt.Println(text)

	// message := greetings.Hello("Pavel")
	// fmt.Println(message)

	// мой цвет
	var favoriteColor = "white"
	fmt.Println("Мой любимый цвет", favoriteColor)

	// мой год рождения и возраст
	birthYear, age := 1995, 29
	fmt.Println("Я родился в", birthYear, "мне", age)

	// блочное присваивание
	var (
		firstName = "pavel"
		lastName  = "Lysov"
	)
	fmt.Println("Я", firstName, lastName)

	var ageInDays int
	ageInDays = 365 * age
	fmt.Println("мой возраст в днях", ageInDays)
}
