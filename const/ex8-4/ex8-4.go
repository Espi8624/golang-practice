package main

import "fmt"

/** 상수는 좌변으로 사용할 수 없다
 */

const PI = 3.14
const FloatPI float64 = 3.14

func main() {
	var a int = PI * 100
	// 타입 차이로 인해 오류 발생
	// var b int = FloatPI * 100

	fmt.Println(a)
	// fmt.Println(b)
}
