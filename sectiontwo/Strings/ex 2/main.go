package main

import (
	"fmt"
	"strings"
)

func main() {
	m1 := "Hello word"
	m2 := "word"
	fmt.Println(strings.Split(m1, " "), m1+m2)
}

/* func main() {
	m1 := "Hello word"
	//m2 := "word"
	fmt.Println(strings.Split(m1, " "))
} */
