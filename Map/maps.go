package main

import (
	"fmt"
)

func main() {
	var nameMap map[int]string
	nameMap = map[int]string{1: "韩", 2: "王", 3: "李"}
	fmt.Println(nameMap)
	jobNameMap := make(map[string]string)
	jobNameMap["李"] = "go开发"
	jobNameMap["王"] = "java开发"
	jobNameMap["韩"] = "测试"
	fmt.Println(jobNameMap)
	fmt.Println(jobNameMap["王"])
	fmt.Println(jobNameMap["韩"])
	fmt.Println(jobNameMap["李"])

	map1 := make(map[string][]string)

	map1["研发部"] = append(map1["研发部"], "weplay", "贪吃蛇", "会玩")
	map1["运营部"] = append(map1["运营部"], "广告运营", "视频运营", "宣传")
	value, exist := map1["研发部"]
	fmt.Println(value, exist)
	fmt.Println(map1["研发部"])

	for k, v := range map1 {
		fmt.Println("k = ", k, "v=", v)
	}
	delete(map1, "研发部")
	for k, v := range map1 {
		fmt.Println("k = ", k, "v=", v)
	}
}
