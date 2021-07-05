package controller

import (
	"con-system/logic"
	"con-system/models"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// SignupHandler 处理注册请求的函数
func SignupHandler(c *gin.Context) {
	// 1.获取参数和参数校验
	p := new(models.ParamSignup)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		// 记录日志
		zap.L().Error("SignupHandler with validate params(ShouldBindJSON)", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	//errLenPass := len(p.Username) == 0 || len(p.Password) == 0
	//errPwdPass := p.Password != p.RePassword
	//if errLenPass || errPwdPass {
	//	zap.L().Error("SignupHandler with validate params(validate)")
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": 1,
	//		"msg":  "请求参数错误",
	//	})
	//	return
	//}
	zap.L().Info("请求数据", zap.String("username", p.Username), zap.String("password", p.Password), zap.String("re_password", p.RePassword))
	fmt.Println("p", p)
	// 2.业务处理
	logic.SignUp(p)

	// 3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})

}
