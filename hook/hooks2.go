package main

import (
	"fmt"
	"time"
)

// 基础函数
func add2(a, b int) int {
	return a + b
}

// 钩子函数接口
type HookFunc2 func(a, b int) int

// 日志记录函数
func logHook2(f HookFunc2) HookFunc2 {
	return func(a, b int) int {
		start := time.Now()
		fmt.Printf("开始计算: %d + %d\n", a, b)
		result := f(a, b)
		fmt.Printf("计算结果: %d，耗时: %v\n", result, time.Since(start))
		return result
	}
}

// 性能监控函数
func monitorHook2(f HookFunc2) HookFunc2 {
	return func(a, b int) int {
		start := time.Now()
		result := f(a, b)
		fmt.Printf("性能监控: 耗时: %v\n", time.Since(start))
		return result
	}
}

func main() {
	// 组合使用多个钩子函数
	addWithHooks := monitorHook2(logHook2(add2))
	result := addWithHooks(3, 4)
	fmt.Println("最终结果:", result)
}
