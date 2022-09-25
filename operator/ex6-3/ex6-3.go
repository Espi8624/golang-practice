package main

import "fmt"

func main() {
	var x int8 = 4
	var y int8 = 64

	// %08b : 빈자리에 0으로 채우고 8자리 이진수로 표현
	fmt.Printf("x: %08b, x<<2: %08b, x<<2: %d\n", x, x<<2, x<<2)
	fmt.Printf("y: %08b, y<<2: %08b, y<<2: %d\n", y, y<<2, y<<2)
}
