package controllers

import (
	"con-system/common"
	"con-system/datamodels"
	"con-system/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"strconv"
)

type ProductController struct {
	Ctx iris.Context
	ProductService services.IProductService
}

func (p ProductController) GetAll() mvc.View {
	productArray, _ := p.ProductService.GetProduct()
	return mvc.View{
		Name: "product/view.html",
		Data: iris.Map{
			"productArray": productArray,
		},
	}
}

func (p *ProductController) PostUpdate() {
	product := &datamodels.Product{}
	p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName: "cons"})
	if err := dec.Decode(p.Ctx.Request().Form, product); err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	err := p.ProductService.UpdateProduct(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}

func (p *ProductController) PostAdd() {
	product :=&datamodels.Product{}
	p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName:"cons"})
	if err:= dec.Decode(p.Ctx.Request().Form,product);err!=nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	_,err:=p.ProductService.InsertProduct(product)
	if err !=nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}

func (p *ProductController) GetManager() mvc.View {
	idString := p.Ctx.URLParam("id")
	id,err :=strconv.ParseInt(idString,10,16)
	if err !=nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	product,err:=p.ProductService.GetProductByID(id)
	if err !=nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	return mvc.View{
		Name:"product/manager.html",
		Data:iris.Map{
			"product":product,
		},
	}
}

