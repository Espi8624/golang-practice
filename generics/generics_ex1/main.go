package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// 제한자 정의
// type Integer interface {
// 	int | int8 | int16 | int32 | int64
// }
// type Float interface {
// 	float32 | float64
// }
// type Numeric interface {
// 	Integer | Float
// }

// 일반적으로 [T any, U any, V any] 로 선언
// any, comparable, constraints.Ordered
// func 함수명[(이름) (제한자)](a 이름) {}

//	func print[T any](a T) {
//		fmt.Println(a)
//	}
func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func print(a interface{}) {
	fmt.Println(a)
}

type MyInt int

func main() {
	var a int = 10
	var b int = 20
	fmt.Println(min(a, b))

	var c int16 = 10
	var d int16 = 20
	fmt.Println(min(c, d))

	var e MyInt = 10
	var f MyInt = 20
	fmt.Println(min(e, f))

	// var e float32 = 3.14
	// var f float32 = 3.98
	// fmt.Println(min(e, f))

	// var g string = "Hello"
	// var h string = "World"
	// fmt.Println(min(g, h))
}
