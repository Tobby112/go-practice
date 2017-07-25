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

	j := `[{"First":"James","Last":"Bond","Age":30},{"First":"Tobby","Last":"Lai","Age":20}]`
	fmt.Printf("json:%s\n", j)

	xp := []person{}
	// use package "encoding/json"
	err := json.Unmarshal([]byte(j), &xp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("xp:%+v\n", xp)

	for i, v := range xp {
		fmt.Println(i, v)
		fmt.Println("\t", v.First, v.Last, v.Age)
	}
}
