package main

import "fmt"

func main() {
	var num = 125
	fmt.Printf("num:%d %b %X %q %p\n", num, num, num, num, &num)
}
