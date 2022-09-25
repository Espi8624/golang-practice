package main

import (
	"ex16_practice/usepkg/custompkg"
	"ex16_practice/usepkg/exinit"
	// "github.com/guptarohit/asciigraph"
	// "github.com/tuckersGo/musthaveGo/ch16/expkg"
)

/*
* cd usepkg
* go mod init ex16_practice/usepkg
* go mod tidy <- 현재 사용하고 있는 패키지 다운
* cd program
* go build <- exe 파일 생성
* .\program.exe <- exe 파일 실행
 */

func main() {
	custompkg.PrintCustom()
	// 앞글자가 소문자면 공개되지 않음
	// custompkg.printcustom2()

	// s := custompkg.Student{}
	// s.Name = "Tom"
	// s.Age = 32
	// fmt.Println(s.Name, s.Age)
	// fmt.Println(custompkg.PI)
	// custompkg.Ratio = 10
	// fmt.Println(custompkg.Ratio)
	// expkg.PrintSample()

	// data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	// graph := asciigraph.Plot(data)

	// fmt.Println(graph)

	exinit.PrintD()
}
