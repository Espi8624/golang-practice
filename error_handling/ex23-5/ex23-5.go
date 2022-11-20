package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		// panic 은 어디서 문제가 생겼는지 바로 확인가능
		panic("b는 0일 수 없습니다.")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	divide(9, 3)
	divide(9, 0)
}
