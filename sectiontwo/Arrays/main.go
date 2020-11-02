package main

func swap(m1, m2 *int) {
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp
}

//import "fmt"

//arrays

// func todo() {
//ar arr []int
//	arr := []int{1, 2, 3, 4}
//	fmt.Println(arr)
//}

//func main() {
//	todo()

//fmt.Println()
//}
