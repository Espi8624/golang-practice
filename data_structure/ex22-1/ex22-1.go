package main

import (
	"container/list"
	"fmt"
)

func main() {
	// list 생성
	v := list.New()
	// 4라는 요소 삽입
	e4 := v.PushBack(4)
	e1 := v.PushFront(1)
	// 3을 e4전에 넣어라
	v.InsertBefore(3, e4)
	// 2를 e1 후에 넣어라
	v.InsertAfter(2, e1)

	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}

	fmt.Println()
	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
}
