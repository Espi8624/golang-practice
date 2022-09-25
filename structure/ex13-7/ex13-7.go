package main

import (
	"fmt"
	"unsafe"
)

// type User struct {
// 	// padding 때문에 40byte 를 사용
// 	A int8 // 1
// 	B int  // 8
// 	C int8 // 1
// 	D int  // 8
// 	E int8 // 1
// }

type User struct {
	// byte 정렬을 잘 하면 메모리를 아낄 수 있음
	A int8 // 1
	C int8 // 1
	E int8 // 1
	B int  // 8
	D int  // 8
}

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}
