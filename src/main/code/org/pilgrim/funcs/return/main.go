package main

import (
	"errors"
	"fmt"
)

func main(){
	fmt.Println("hello functions with return :)")
	port := 3000
	isStrarted := startWebService(port, 2)
	started := "no"
	if(isStrarted){
		started = "yes"
	}
	fmt.Println("Is service started?", started)

	result := startWebService2(port, 3)
	fmt.Println(result)

	p, err := startWebService3(port, 3)
	fmt.Println(p, err)
}

func startWebService(p int, numberOfRetries int) bool{
	fmt.Println("Starting service...")
	// do some implementation

	fmt.Println("Service started on port:", p)
	fmt.Println("Number of retries:", numberOfRetries);

	return true
}

func startWebService2(p int, numberOfRetries int) error{
	fmt.Println("Starting service...")
	// do some implementation

	return errors.New("Service cannot be started")
}


func startWebService3(p int, numberOfRetries int) (int, error){
	fmt.Println("Starting service...")
	// do some implementation

	return p, errors.New("Service cannot be started")
}