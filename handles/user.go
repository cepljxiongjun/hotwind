package handles

import (
	"fmt"

	"github.com/cepljxiongjun/hotwind/helper"
	"github.com/cepljxiongjun/hotwind/service"
	"github.com/gin-gonic/gin"
)

// User all func
var User user

type user struct {
}

//  Create a new user
func (user) Create(c *gin.Context) {
	username := c.PostForm("username")
	passowrd := c.PostForm("password")
	fmt.Println(len(passowrd), passowrd, username)
	if len(passowrd) < 6 || len(passowrd) > 20 {
		c.JSON(400, gin.H{
			"message": "密码长度为6-20",
		})
		return
	}
	if err := service.User.Create(username, passowrd); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func (user) Login(c *gin.Context) {
	username := c.PostForm("username")
	passowrd := c.PostForm("password")
	u, err := service.User.GetByUsername(username)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "未找到该用户",
		})
		return
	}
	bool := helper.CheckPassword(u.PasswordHash, passowrd)
	if !bool {
		c.JSON(400, gin.H{
			"message": "密码错误",
		})
		return
	}
	jwt, _ := helper.JwtEncode(u)
	c.JSON(200, gin.H{
		"message": "登录成功",
		"jwt":     jwt,
	})
	return
}
