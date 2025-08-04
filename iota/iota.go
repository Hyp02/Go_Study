package main

import (
	"fmt"
	"time"
)

func sayHello() {

	fmt.Println("Hello")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second) // 等待 goroutine 完成
	}

}
func arr() {
	arr := [3]int{0, 1, 2}
	arr1 := []int{0, 1, 2}
	arr2 := [...]int{0, 1, 2}
	var arr11 [3]int = [3]int{0, 1, 2}
	var arr3 = []int{1, 2, 3}
	fmt.Println("Array len: ", len(arr))
	fmt.Println("Array len: ", len(arr1))
	fmt.Println("Array len: ", len(arr2))
	for _, i := range arr {
		fmt.Println(i)
	}
}

func main() {
	arr()
	// go sayHello() // 在 goroutine 中异步执行
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("主线程 %d\n", i)
	// 	time.Sleep(1 * time.Second) // 等待 goroutine 完成
	// }
}

// func main() {
// 	var s user
// 	s = 1
// 	fmt.Println(s)
// 	fmt.Println(one)
// 	fmt.Println(two)
// 	fmt.Println(three)
// 	fmt.Println(four)
// 	defer fmt.Println("defer")
// 	return
// }
