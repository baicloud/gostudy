package user

type User struct {
	Username  string
	Sex       string
	Age       int
	AvatarUrl string
}

//go语言没有构造函数，需要自己定义
func NewUser(username string,sex string,age int,avatarurl string) *User{
	var u *User=new(User)
	u.Username=username
	u.Sex=sex
	u.Age=age
	u.AvatarUrl=avatarurl
	return u
}