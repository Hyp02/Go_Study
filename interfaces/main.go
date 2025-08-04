package main

import "fmt"

// 定义接口
type Speaker interface {
	Speak() string
}

// 定义结构体
type Dog struct {
	Name string
}

// 实现接口方法
func (d Dog) Speak() string {
	return fmt.Sprintf("%s says: Woof!", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("%s says: Meow!", c.Name)
}

func main() {
	// 创建接口变量
	var speaker Speaker

	// Dog实现了Speaker接口
	dog := Dog{Name: "Buddy"}
	speaker = dog
	fmt.Println(speaker.Speak())
	dogSperaker, ok := speaker.(Dog)
	fmt.Println(dogSperaker.Name, ok)
	fmt.Println(speaker.Speak())

	// Cat实现了Speaker接口
	cat := Cat{Name: "Whiskers"}
	speaker = cat
	fmt.Println(speaker.Speak())

	// 接口类型断言
	if dogSpeaker, ok := speaker.(Dog); ok {
		fmt.Println("It's a dog:", dogSpeaker.Name)
	} else {
		fmt.Println("Not a dog")
	}
}
