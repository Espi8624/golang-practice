package main

import "fmt"

func main() {
	var str string = "Hello world"
	// 메모리를 새로 할당받아 사용
	var slice []byte = []byte(str)

	slice[2] = 'a'

	fmt.Println(str)
	fmt.Printf("%s\n", slice)
}
