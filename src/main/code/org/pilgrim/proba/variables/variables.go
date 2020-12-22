package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	// 1. /////////////////////////
	var i int
	i = 42
	fmt.Println(i)

	// 2. /////////////////////////
	var f32 float32 = 3.14
	fmt.Println(f32)

	// 3. /////////////////////////
	firstName := "Arthur"
	fmt.Println(firstName)

	// 4. /////////////////////////
	co := complex(3, 4)
	fmt.Println(co)

	// 5. /////////////////////////
	r, im := real(co), imag(co)
	fmt.Println(r, im)

	// 6. /////////////////////////
	var fn *string
	fmt.Println(fn)

	fn = new(string)
	*fn = "Bob"
	fmt.Println(fn)
	fmt.Println(*fn)

	// 7. /////////////////////////
	fy := "Tom"
	fmt.Println(fy)

	ptr := &fy
	fmt.Println(ptr, *ptr)

	fy = "Lee"
	fmt.Println(ptr, *ptr)
}
