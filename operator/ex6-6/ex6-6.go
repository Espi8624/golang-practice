package main

import "fmt"

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	// 0.3을 정확하게 표현하지 못하기 때문에 가장 가까운 수로 표현
	// 실제로 정확하게 같지가 않음
	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c)
	fmt.Printf("%v\n", a+b)
}
