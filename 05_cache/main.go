package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

// testing struct for go-cache
type student struct {
	id   int
	name string
}

// just print out the input string
func MyFunction(s string) {
	fmt.Printf("MyFunction:%s\n", s)
}

func main() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", cache.DefaultExpiration)

	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("baz", 42, cache.NoExpiration)

	// Since Go is statically typed, and cache values can be anything, type
	// assertion is needed when values are being passed to functions that don't
	// take arbitrary types, (i.e. interface{}). The simplest way to do this for
	// values which will only be used once--e.g. for passing to another
	// function--is:
	if x, found := c.Get("foo"); found {
		foo := x.(string)
		MyFunction(foo)
	}

	// ...
	// foo can then be passed around freely as a string
	// Want performance? Store pointers!
	sudentTobby := student{99999, "TOBBY"}
	c.Set("gogo", &sudentTobby, cache.DefaultExpiration)
	if x, found := c.Get("gogo"); found {
		foo := x.(*student)
		fmt.Printf("student:%v\n", foo)
	}

	count := c.ItemCount()
	fmt.Printf("cache count:%d\n", count)
	c_map := c.Items()
	for key, value := range c_map {
		fmt.Printf("%s : %v\n", key, value)
	}
}
