package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}

	// 불필요한 메모리 할당 없는 방식
	idx := 2
	slice = append(slice, 0)
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = 100

	// idx := 2
	// slice = append(slice[:idx], append([]int{100}, slice[idx:]...)...)

	// slice = append(slice, 0)
	// idx := 2
	// for i := len(slice) - 2; i >= idx; i-- {
	// 	slice[i+1] = slice[i]
	// }
	// slice[idx] = 100

	fmt.Println(slice)
}
