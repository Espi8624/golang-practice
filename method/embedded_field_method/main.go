package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// Sprintf => 문자열을 생성하여 반환
func (u User) String() string {
	return fmt.Sprintf("%s, %d", u.Name, u.Age)
}

type VIPUser struct {
	User     // embedded field, 상속 x, 포함 o
	VIPLevel int
}

func (v VIPUser) vipLevel() int {
	return v.VIPLevel
}

func main() {
	vip := VIPUser{User{"Hwang", 34}, 5}
	fmt.Println(vip.String())
	fmt.Println(vip.vipLevel())
}
