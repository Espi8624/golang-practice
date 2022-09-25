package main

import "fmt"

func changeArray(array2 [5]int) {
	array2[2] = 200
}

func changeSlice(slice2 []int) {
	slice2[2] = 200
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	// 새로운 공간으로 복사되었기 때문에 값이 변하지 않음.
	changeArray(array)
	changeSlice(slice)

	fmt.Println(array)
	fmt.Println(slice)
}
