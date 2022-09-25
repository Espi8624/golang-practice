package custompkg

import "fmt"

type Student struct {
	Name  string
	Age   int
	score int
}

// 외부로 공개
var Ratio int

const PI = 3.14

// 외부로 비공개
var ttt int

const pI2 = 3.1415

func PrintCustom() {
	fmt.Println("This is custom package!")
}

func printcustom2() {
	fmt.Println("This is custom package2!")
}
