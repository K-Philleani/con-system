package models

// ParamSignup 定义请求参数的结构体
type ParamSignup struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 定义登录请求的结构体
type ParamLogin struct {
	Username string `json:"username" bing:"required"`
	Password string `json:"password" bing:"required"`
}
