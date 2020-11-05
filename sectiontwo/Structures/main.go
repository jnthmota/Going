package main

import (
	"fmt"
)

//structure encapsulation
type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func main() {
	//c := Car{"Posche", 1, 2}
	// another way
	c := Car{
		Name:    "Posche",
		Age:     1,
		ModelNo: 2,
	}
	//var c1 Car
	fmt.Println(c)
}
