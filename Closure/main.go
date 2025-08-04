package main

import (
	"fmt"
)

func main() {
	i := 0
	count := counts(i)
	for i := 0; i < 10; i++ {

		fmt.Println("count():", count())
		i = count()
		fmt.Println("i=", i)
	}
	fmt.Println(i)

	// 准备一个字符串
	str := "hello world"

	// 创建一个匿名函数
	foo := func() {

		// 匿名函数中访问str
		str = "hello dude"
	}
	// 调用匿名函数
	foo()
	fmt.Println(str)

	fii := func() {
		i = 1
	}
	fii()
	fmt.Println(i)
}
func counts(i int) func() int {
	return func() int {
		i++
		return i
	}
}
