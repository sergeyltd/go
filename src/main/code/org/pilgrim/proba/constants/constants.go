package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	// 1. //////////////////////////////////
	const pi = 3.1415
	fmt.Println(pi)

	// 2. //////////////////////////////////
	const c = 3
	fmt.Println(c + 2)
	fmt.Println(c + 1.2)

	// 3. //////////////////////////////////
	const t int = 3
	fmt.Println(t + 2)
	fmt.Println(t + c)
    // fmt.Println(t + 1.2) - wrong
    fmt.Println(float32(t) + 1.2)
}
