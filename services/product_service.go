package services

import (
	"con-system/datamodels"
	"con-system/repositories"
)

type IProductService interface {
	GetProductByID(int64) (*datamodels.Product, error)
	GetProduct() ([]*datamodels.Product, error)
	DeleteProductByID(int64) bool
	InsertProduct(*datamodels.Product) (int64, error)
	UpdateProduct(*datamodels.Product) error
}

type productService struct {
	productRepository repositories.IProduct
}

func NewProductService(repository repositories.IProduct) IProductService{
	return &productService{productRepository:repository }
}

func (p *productService) GetProductByID(productID int64) (*datamodels.Product, error) {
	return p.productRepository.Select(productID)
}

func (p *productService) GetProduct() ([]*datamodels.Product, error) {
	return p.productRepository.SelectAll()
}

func (p *productService) DeleteProductByID(productID int64) bool {
	return p.productRepository.Delete(productID)
}

func (p *productService) InsertProduct(product *datamodels.Product) (int64, error) {
	return p.productRepository.Insert(product)
}

func (p *productService) UpdateProduct(product *datamodels.Product) error {
	return p.productRepository.Update(product)
}