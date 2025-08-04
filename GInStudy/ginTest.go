package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	name string `json:"username"`
	pwd  string `json:"pwd"`
}

func main() {
	r := gin.Default()
	// get请求，路径为/ 处理函数为helloWord
	r.GET("/", HelloWord)
	r.POST("/post", HelloWorldPost)
	r.Run(":8080")
}
func HelloWord(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
		"code":    200,
	})
}
func HelloWorldPost(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	if user.name == "han" && user.pwd == "123" {
		c.JSON(200, gin.H{
			"message":  "登陆成功",
			"username": user.name,
			"pwd":      user.pwd,
			"code":     200,
		})
	} else {
		c.JSON(300, gin.H{"message": "用户名或者密码错误"})
	}
	return
}

// func getUserInfo(c *gin.Context) (string, string) {
// 	name := c.PostForm("username")
// 这是hyp3分支
// }
