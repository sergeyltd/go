package main

import (
	"fmt"
)

func main() {
	type user struct{
		id int
		firstName string
		lastName string
	}

	var u user
	u.id = 1
	u.firstName = "Tom"
	u.lastName = "Riddle"
	fmt.Println(u)

	u2 := user{
		id: 2,
		firstName: "Harry",
		lastName: "Potter",
	}

	fmt.Println(u2)
}
