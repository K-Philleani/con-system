package models

// ParamSignup 定义请求参数的结构体
type ParamSignup struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}
