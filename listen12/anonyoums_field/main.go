package main

import "fmt"

type Address struct {
	City     string
	Province string
}

type User struct {
	Username  string
	Sex       string
	Age       int
	AvatarUrl string
	// 匿名字段默认采用类型名作为字段名
	int
	string
	address Address
}

func main() {
	var user *User = &User{
		Username: "stu01",
		int:      1,
		address: Address{
			City:     "shanghai",
			Province: "shanghai",
		},
	}
	fmt.Println( user.Username,user.address.City)
}
