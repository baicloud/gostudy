package main

import (
	"fmt"
	"github.com/baicloud/gostudy/listen12/user"
)

func main() {
	var u user.User
	u.Username = "stu01"
	u.Sex = "男"
	u.Age = 21
	u.AvatarUrl = "http://text/1.png"
	fmt.Println(u)

	u1 := user.NewUser("stu01", "男", 16, "https://example.com/1.png")
	fmt.Println(u1)
	fmt.Printf("%#v\n", u1)
}
