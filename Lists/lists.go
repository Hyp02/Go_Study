package main

import (
	"container/list"
	"fmt"
)

func main() {

	myList := list.New()
	// 插入元素
	myList.PushBack(1)
	myList.PushBack(2)
	fmt.Println(myList.Front().Value)
	// 删除元素

}
