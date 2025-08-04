package main

import "fmt"

// 1. 定义接口
type Speaker interface {
	Speak() string
}

// 2. 定义函数类型
// SpeakerFunc是一个函数类型，它匹配Speaker接口的Speak方法签名
type SpeakerFunc func() string

// 3. 为函数类型实现接口方法
// 这使得SpeakerFunc类型满足Speaker接口
func (f SpeakerFunc) Speak() string {
	// 直接调用函数本身
	fmt.Println("函数实现接口")
	return f()
}

func main() {
	// 4. 创建普通函数
	sayHello := func() string {
		return "Hello from function value!"
	}

	// 5. 将函数转换为SpeakerFunc类型
	speakerFunc := (SpeakerFunc)(sayHello)

	// 6. 赋值给接口变量
	var speaker Speaker = speakerFunc

	// 7. 通过接口调用方法
	fmt.Println(speaker.Speak())

	// 8. 直接使用匿名函数实现接口
	var s2 Speaker = SpeakerFunc(func() string {
		return "Anonymous function implementation"
	})
	fmt.Println(s2.Speak())
}
