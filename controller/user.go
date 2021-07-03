package controller

import (
	"con-system/logic"
	"con-system/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SignupHandler 处理注册请求的函数
func SignupHandler(c *gin.Context) {
	// 1.获取参数和参数校验
	var p models.ParamSignup
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "请求参数错误",
		})
		return
	}
	// 2.业务处理
	logic.SignUp()
	// 3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})

}
