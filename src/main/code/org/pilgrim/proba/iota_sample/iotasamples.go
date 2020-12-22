package main

import (
	"fmt"
)

const pi = 3.14
const (
	first  = 1
	second = "second"
)
const (
	a = iota
	b = iota
	c = iota + 5
	d = 2 << iota
	e
	f
)

const (
	g = iota
	h
)

func main() {
	fmt.Println(pi)
	fmt.Println(first, second)
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(g, h)

}
