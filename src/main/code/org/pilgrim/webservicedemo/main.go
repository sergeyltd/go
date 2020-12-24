package main

import (
	"github.com/pluralsight/webservicedemo/modules"
	"fmt"
)

func main(){
	u := modules.User{
		ID: 3,
		FirstName: "Tom",
		LastName: "Riddle",
	}

	u.ID = 4
	fmt.Println("Hello Webservice Demo")
	fmt.Println(u)
}