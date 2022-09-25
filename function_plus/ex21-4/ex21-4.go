package main

import "fmt"

type OpFn func(int, int) int

func getOperator(op string) OpFn {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

// 함수 리터럴 => 내부상태(State)를 가질 수 있음.
func main() {
	var operator func(int, int) int
	operator = getOperator("*")

	var result = operator(3, 4)
	fmt.Println(result)
}
