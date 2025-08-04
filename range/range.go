package main

import (
	"fmt"
)

func main() {
	var a = []int{}
	b := make([]int, 0, 3)
	var slice []int
	res1 := forFun(a, 10, "a切片")
	res2 := forFun(b, 10, "b切片")
	res3 := forFun(slice, 10, "slice切片")
	fmt.Println("=======a==========")
	for i, v := range res1 {
		fmt.Println("a index = ", i, "value=", v)
	}
	fmt.Println("=======b==========")

	for i, v := range res2 {
		fmt.Println("b index = ", i, "value=", v)
	}
	fmt.Println("=======c==========")

	for i, v := range res3 {
		fmt.Println("s index = ", i, "value=", v)
	}

}
func forFun(a []int, lens int, s string) (res []int) {
	for i := 0; i < lens; i++ {
		a = append(a, i)
		fmt.Println(s, a, "容量：", cap(a), "长度", len(a))
	}
	return a

}
