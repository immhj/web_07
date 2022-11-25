package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	dao "web_06/dao/mysql"
	dao2 "web_06/dao/redis"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if dao.Selectuser(username) == false {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "用户名已注册",
		})
		return
	}
	dao.Createuser(username, password)
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	if dao.Selectuser(username) == true {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "用户名不存在",
		})
		return
	}
	password := c.PostForm("password")
	if dao.Selectpasswordbyusername(username) != password {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "密码输入错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "登陆成功",
	})
}

func Setquestion(c *gin.Context) {
	school := c.PostForm("school")
	favor := c.PostForm("favor")
	food := c.PostForm("food")
	username := c.PostForm("username")
	if dao.Selectuser(username) == true {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "用户名不存在",
		})
	}
	dao.Setquestion(food, favor, school, username)
	c.JSON(http.StatusOK, gin.H{
		"message": "set correct",
	})
}

func Findpassword(c *gin.Context) {
	school := c.PostForm("school")
	favor := c.PostForm("favor")
	food := c.PostForm("food")
	username := c.PostForm("username")
	if dao.Findpassword(food, favor, school, username) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "密保信息错误",
		})
	}
	password := dao.Selectpasswordbyusername(username)
	c.JSON(http.StatusOK, gin.H{
		username: password,
	})
}

func Info(c *gin.Context) {
	infomation := c.PostForm("info")
	dao.Info(infomation)
	c.JSON(http.StatusOK, gin.H{
		"message": "设置成功",
	})
}

func UPP(c *gin.Context) {
	username := c.PostForm("username")
	flag := dao2.Upout(username)
	if flag == 1 {
		c.JSON(http.StatusOK, gin.H{
			"message": "只可以点一次赞",
		})
		return
	}
	dao2.Upin(username)
	c.JSON(http.StatusOK, gin.H{
		"message": "感谢您的点赞",
	})
}
