package main

import (
	"fmt"
	"strconv"
)

// 인터페이스
type Stringer interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
	String() string
}

func PrintMin[T Stringer](a, b T) {
	if a < b {
		fmt.Println(a.String())
	} else {
		fmt.Println(b.String())
	}
}

// 타입제한자
// type Integer interface {
// 	~int8 | ~int16 | ~int32 | ~int64 | ~int
// }

// func Print1(a Stringer) {
// 	fmt.Println(a.String())
// }

// func Print2[T Stringer](a T) {
// 	fmt.Println(a.String())
// }

type MyInt int

func (m MyInt) String() string {
	return strconv.Itoa(int(m))
}

type MyString struct {
	name string
}

func (m MyString) String() string {
	return m.name
}

func main() {
	var m1 MyInt = 10
	var m2 MyInt = 20
	PrintMin(m1, m2)

	// m := MyString{"Name"}
	// Print1(m)
	// Print2(m)
}
