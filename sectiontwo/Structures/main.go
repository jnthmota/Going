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
	c := Car{"Posche", 1, 2}
	//var c1 Car
	fmt.Println(c)
}
