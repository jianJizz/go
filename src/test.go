package main


import "fmt"
var (
	a int
	b bool
)

// TODO:学习go的基本语法
func main() {
	// var i int
	// var f float64
	// var b bool
	// var s string
	// fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// var x = true
	// fmt.Println(x)

	// y := 1.5 //这是一个声明语句左侧的变量不应该是一件声名过的，否则会编译错误 且这种形式只能在函数体中出现

	// x, y, z := 1, "a", false

	// a = 10
	// b = true
	// fmt.Println(x, y, z)
	// fmt.Println(a, b)

	_, numbers, str := numbers()
	fmt.Println(numbers, str)
}

func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}
