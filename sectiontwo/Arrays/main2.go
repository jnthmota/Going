package main

import "fmt"

//arrays

func todo() {
	//ar arr []int
	arr := []int{1, 2, 3, 4}
	arr2 := []string{"Hello", "Word", "I'm"}

	arr2 = append(arr2, "Jonathan")
	fmt.Println(arr, arr2)
}

func main() {
	todo()
	//fmt.Println()
}
