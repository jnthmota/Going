package main

import (
	"fmt"
	"strings"
)

func main() {
	m1 := "Hello word"
	m2 := "word"
	fmt.Println(strings.Contains(m1, m2))
}
