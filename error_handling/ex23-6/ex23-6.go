package main

import "fmt"

func f() {
	fmt.Println("f() 함수 시작")
	defer func() {
		// 왠만하면 안쓰는걸 추천
		// 배포가 되었을 때 죽으면 안될때 사용
		if r := recover(); r != nil {
			fmt.Println("panic 복구 - ", r)
		}
	}()

	g()
	fmt.Println("f() 함수 끝")
}

func g() {
	fmt.Printf("9/3 = %d\n", h(9, 3))
	fmt.Printf("9/0 = %d\n", h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic("제수는 0일 수 없습니다.")
	}
	return a / b
}

func main() {
	f()
	fmt.Println("프로그램이 계속 실행됨")
}
