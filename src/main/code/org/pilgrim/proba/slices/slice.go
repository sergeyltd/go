package main

import (
	"fmt"
)

func main() {
	// 1. ////////////////////////////
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	slice := arr[:]
	fmt.Println(arr, slice)

	arr[1] = 42
	slice[2] = 27
	fmt.Println(arr, slice)

	// 2. ////////////////////////////
	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)

	slice2 = append(slice2, 4)
	fmt.Println(slice2)

	slice2 = append(slice2, 5, 6, 7)
	fmt.Println(slice2)

	s3 := slice2[1:]
	s4 := slice2[:2]
	s5 := slice2[1:4]
	fmt.Println(s3, s4, s5)
}
