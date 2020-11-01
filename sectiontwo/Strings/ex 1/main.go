package main

import (
	"fmt"
	"strings"
)

// replace na minha string
func main() {
	m1 := "Hello word"
	//m2 := "word"
	fmt.Println(strings.ReplaceAll(m1, "Hello", "My"))
}
