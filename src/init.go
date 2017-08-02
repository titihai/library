package src

import "github.com/gin-gonic/gin"
import "library/src/handler"

/*
项目的一些初始化工作
*/
func init() {
	r := gin.Default()
	r.GET("/register", handler.UserRegister)
	r.Run() // listen and serve on 0.0.0.0:8080
}
