package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"fname"`
	Last  string
	Age   int
}

func main() {
	// if use json tags in definition of fields
	// the member name will become useless when doing unmarshal
	// unmarshal just takes the tags to decoding the json string
	j := `[{"fname":"James","Last":"Bond","Age":30},{"fname":"Tobby","Last":"Lai","Age":20}]`
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
