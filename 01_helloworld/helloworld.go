package main

import "fmt"

func increaseTwo() func() int {
	z := 0
	return func() int {
		z += 2
		return z
	}
}

func main() {
	fmt.Println("hello world!", y)
	x := 0
	// closure
	increasement := func() int {
		x++
		return x
	}
	fmt.Printf("x=%d\n", increasement())

	// closure
	increasement = func() int {
		x -= 2
		return x
	}
	fmt.Printf("x=%d\n", increasement())

	increasement = increaseTwo()
	fmt.Printf("z=%d\n", increasement())
}

var y = 123
