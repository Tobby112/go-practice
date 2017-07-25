package main

import "fmt"

type test struct {
	a int
	b int
}

func (t test) multiply() int {
	fmt.Printf("[multiply] t at %p\n", &t)
	t.a = 100
	return t.a * t.b
}

func (t *test) multiplyPtr() int {
	fmt.Printf("[multiplyPtr] t at %p\n", t)
	t.b = 50
	return t.a * t.b
}

func main() {
	var r int
	t1 := test{5, 6}
	r = t1.multiply()
	fmt.Printf("[main] t at %p\n", &t1)
	fmt.Println("t1: ", t1)
	fmt.Println("r: ", r)
	r = (&t1).multiplyPtr()
	fmt.Println("r: ", r)
	fmt.Println("t1: ", t1)

	var t2 = &t1
	fmt.Printf("[main] t2 at %p\n", t2)
	t2.a = 30
	r = (&t1).multiplyPtr()
	fmt.Println("r: ", r)
}
