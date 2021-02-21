package main

import (
	"con-system/common"
	"con-system/datamodels"
	"con-system/repositories"
	"con-system/services"
	"github.com/kataras/iris/v12"
	"log"
)

func main() {
	// 1.创建iris实例
	app := iris.Default()
	// 2.设置错误模式， 在MVC模式下提示错误
	db, err := common.NewMysqlConn()
	if err != nil {
		log.Println("data err", err)
		return
	}
	db.Ping()
	log.Println("数据库连接成功")
	app.Get("/product/add", func(ctx iris.Context) {
		productRepository := repositories.NewProductManager("product", db)
		productService := services.NewProductService(productRepository)
		n, err := productService.InsertProduct(&datamodels.Product{
			ID:           1,
			ProductName:  "可乐",
			ProductNum:   100,
			ProductImage: "",
			ProductUrl:   "",
		})
		if err != nil {
			ctx.JSON(iris.Map{
				"msg": "failed",
			})
			panic(err)
		}
		ctx.JSON(iris.Map{
			"msg": "success",
			"id": n,
		})
	})

	app.Run(iris.Addr("localhost:8080"))
}
