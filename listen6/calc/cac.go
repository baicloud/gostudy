package calc

//开头字母小写只在包内可见，大写可被其他包导入
var a int = 10
var A int = 11

func Add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
