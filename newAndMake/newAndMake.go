package main

import (
	"fmt"
)

func main() {
	var i *int
	i = new(int)
	*i = 10
	fmt.Println(i)
	fmt.Println(*i)

	*i = *i + 1
	fmt.Println(&i)
	fmt.Println(*i)
	*i = *i + 1
	fmt.Println(&i)
	fmt.Println(*i)

	ff1, f := namedRetValues()
	fmt.Println(ff1, f)
	var ff2 func()
	ff2 = ff
	ff2()
	nmF := func(x int) int {
		fmt.Println("我是匿名函数")
		return x
	}(100)
	fmt.Println(nmF) // 100
}
func ff() {
	fmt.Println("ff")
}
func namedRetValues() (id int, name string) {

	name = "xx"
	id = 1

	return id, name
}
