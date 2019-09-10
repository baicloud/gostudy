package main

import "fmt"

//定义结构体
type User struct {
	username  string
	sex       string
	age       int
	AvatarUrl string
}

func main() {
	//声明与初始化结构体
	//初始化方法一
	var user User //普通方式声明结构体
	user.username = "stu01"
	user.sex = "男"
	user.age = 18
	user.AvatarUrl = "http://www.example.com/1.png"
	fmt.Println(user)

	//初始化方法二
	var user01 User = User{
		username:  "stu02",
		sex:       "女",
		age:       17,
		AvatarUrl: "http://www.example.com/2.png",
	}
	fmt.Println(user01)
	user02 := User{
		username:  "stu03",
		sex:       "女",
		age:       19,
		AvatarUrl: "http://www.example.com/3.png",
	}
	fmt.Println(user02)
}
