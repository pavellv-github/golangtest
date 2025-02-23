package main

import (
	"fmt"
)

func main() {
	type User struct {
		firstName string
		lastName  string
		age       int
		birthDate string
	}

	dataUser := User{"Pavel", "lysov", 29, "15/07/1995"}
	fmt.Println("im:", dataUser.firstName, dataUser.age, dataUser.lastName, dataUser.birthDate)

	fmt.Println("Hello my age", dataUser.age)

	dataUser.age = 12
	fmt.Println("Hello my age", dataUser.age)

	dataUser.firstName = "asd"
	fmt.Println(dataUser)

	dataGuest := struct {
		firstName     string
		lastName      string
		age, dayVisit int
	}{
		"guest name",
		"",
		22,
		29,
	}

	fmt.Println(dataGuest)
}
