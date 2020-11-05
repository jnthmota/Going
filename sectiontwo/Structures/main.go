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

//1
func (c Car) Print() {
	fmt.Println(c)
}

//2
func (c Car) Drive() {
	fmt.Println("Driving fast ....")
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
	//fmt.Println(c)
	c.Print()
	c.Drive()

}
