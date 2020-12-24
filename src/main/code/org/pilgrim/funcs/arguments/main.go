package main

import "fmt"

func main(){
	fmt.Println("hello functions with agruments :)")
	port := 3000
	startWebService(port, 2)
}

func startWebService(p int, numberOfRetries int){
	fmt.Println("Starting service...")
	// do some implementation

	fmt.Println("Service started on port:", p)
	fmt.Println("Number of retries:", numberOfRetries);
}