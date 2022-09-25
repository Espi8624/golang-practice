package main

import "fmt"

func main() {
	str := "Hello 월드"
	// rune -> int32의 별칭타입
	// []rune -> 한칸에 4byte
	arr := []rune(str)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입: %T, 값: %d, 문자값: %c\n", arr[i], arr[i], arr[i])
	}
}
