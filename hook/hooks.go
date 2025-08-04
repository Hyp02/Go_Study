package main

import (
	"fmt"
	"time"
)

// 基础函数
func add(a, b int) int {
	return a + b
}

// 钩子函数接口
type HookFunc func(a, b int) int

// 日志记录函数
func logHook(f HookFunc) HookFunc {
	return func(a, b int) int {
		start := time.Now()
		fmt.Printf("开始计算: %d + %d\n", a, b)
		result := f(a, b)
		fmt.Printf("计算结果: %d，耗时: %v\n", result, time.Since(start))
		return result
	}
}

// func main() {
// 	// 使用钩子函数
// 	addWithLog := logHook(add)
// 	result := addWithLog(3, 4)
// 	fmt.Println("最终结果:", result)
// }
