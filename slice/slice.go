package main

import (
	"fmt"
)

func main() {
	var i []int
	fmt.Println(i, cap(i), len(i))
	i = append(i, 1)
	fmt.Println(i, cap(i), len(i))
	i = append(i, 2)
	fmt.Println(i)
	i = append(i, 3, 4)
	fmt.Println(i)
	fmt.Println("---------------")
	var s = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s[1:3:5])
	fmt.Println(s, s[:1], len(s), len(s[:1]), cap(s), cap(s[:1]))
	fmt.Println(s, s[1:2], len(s), len(s[1:2]), cap(s), cap(s[1:2]))
	fmt.Println(s, s[2:3], len(s), len(s[2:3]), cap(s), cap(s[2:3]))
	fmt.Println(s, s[3:5], len(s), len(s[5:5]), cap(s), cap(s[3:5]))
	fmt.Println("------append-----------")
	is := append(i, s...)
	fmt.Println(is, cap(is), len(is), cap(is))

}
