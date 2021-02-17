package main

import "github.com/kataras/iris/v12"

func main() {
	// 1.创建iris实例
	app := iris.New()
	// 2.设置错误模式， 在MVC模式下提示错误
	app.Logger().SetLevel("debug")
	// 3.注册模板
	template := iris.HTML("./backend/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	// 4.设置模板目录
	app.HandleDir("/assets","./backend/web/assets")
	// 出现异常跳转到指定的页面
	//app.OnAnyErrorCode(func(ctx iris.Context) {
	//		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问页面出错"))
	//		ctx.ViewLayout("")
	//		ctx.View("shared/error.html")
	//})
	// 5.注册控制器
	// 6.启动服务
	app.Run(iris.Addr("localhost:8080"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)
}
