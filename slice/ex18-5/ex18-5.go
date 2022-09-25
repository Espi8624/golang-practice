package main

import "fmt"

//1
// func addNum(slice *[]int) {
// 	*slice = append(*slice, 4)
// }

// func main() {
// 	slice := []int{1, 2, 3}
// 	addNum(&slice)

// 	fmt.Println(slice)
// }

// 2 => 일반적으로 더 나은 방법
func addNum(slice []int) []int {
	slice = append(slice, 4)
	return slice
}

func main() {
	slice := []int{1, 2, 3}
	slice = addNum(slice)

	fmt.Println(slice)
}
