package handler

import (
	"library/src/util"

	"github.com/gin-gonic/gin"
)

/*
用户相关handler
*/

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	err := util.SendEmail("lianchonghai@qq.com", "<h>Hello,Tiger</h>", "Hello")
	if err != nil {
		util.Logger.Error(err)
	}
	c.String(200, "Send e-mail success")
}
