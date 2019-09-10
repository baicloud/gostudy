package main

import "fmt"

type User struct {
	Username  string
	Sex       string
	Age       int
	AvatarUrl string
}

func main() {
	var user *User = &User{}
	fmt.Println(user)
	var user01 *User = &User{}
	//在struct中，无论使用的是指针的方式声明还是普通方式，访问其成员都使用"."，在访问的时候编译器会自动把 stu2.name 转为 (*stu2).name。
	//(*user01).Username="stu01"
	user01.Username = "stu02"
	fmt.Printf("user01=%#v\n", user01)
	var user02 *User = &User{
		Username:  "stu03",
		Sex:       "男",
		Age:       18,
		AvatarUrl: "http://example.com/1.png",
	}
	fmt.Printf("user02=%v\n", user02)
	var user03 *User = new(User)
	user03.Username = "stu04"
	user03.Sex = "女"
	user03.Age = 18
	fmt.Printf("user03=%v\n", user03)
}
