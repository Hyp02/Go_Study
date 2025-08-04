// 回调函数定义后不会立即执行，在他需要的时候才执行
package main

import (
	"fmt"
	"time"
)

func main() {

	// 定义回调函数
	callback := func(operation string) bool {
		fmt.Println("执行操作:", operation)
		if operation != "查询" {
			return false
		}
		return true
	}
	// 发起数据库查询，异步调用，不影响主进程
	QueryDB("1SELECT * FROM users", callback)
	fmt.Println("查询已发起")
	time.Sleep(2 * time.Second)
}
func QueryDB(sql string, callback func(string) bool) {
	go func() {
		// 模拟数据库查询
		time.Sleep(1 * time.Second)
		operation := "查询"
		b := callback(operation)
		fmt.Println("执行sql为:" + sql)
		if b {
			fmt.Println("查询成功")
		} else {
			fmt.Println("查询失败")
		}
	}()
}
