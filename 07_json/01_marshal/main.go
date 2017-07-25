package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{First: "James", Last: "Bond", Age: 30}
	p2 := person{First: "Tobby", Last: "Lai", Age: 20}
	xp := []person{p1, p2}
	fmt.Printf("go data:%+v\n", xp)

	// use package "encoding/json"
	bs, err := json.Marshal(xp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("json:", string(bs))
}
