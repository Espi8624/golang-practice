package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}

	// 1
	slice2 := make([]int, len(slice1))
	// copy(dst, src)
	// slice1 -> slice2 로 복사
	copy(slice2, slice1)

	// 2
	// slice2 := append([]int{}, slice1...)

	// 3
	// slice2 := make([]int, len(slice1))
	// for i, v := range slice1 {
	// 	slice2[i] = v
	// }

	slice2[1] = 100
	fmt.Println("slice1", slice1)
	fmt.Println("slice1", slice2)
}
