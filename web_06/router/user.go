package router

import (
	"github.com/gin-gonic/gin"
	"web_06/api"
)

func Initrouter() {
	r := gin.Default()
	r.POST("/register", api.Register)         //注册
	r.POST("/login", api.Login)               //登录
	r.POST("/setquestion", api.Setquestion)   //设置密保
	r.POST("/findpassword", api.Findpassword) //通过密保找回密码
	r.POST("/info", api.Info)
	r.POST("/up", api.UPP)
	r.Run()
}
