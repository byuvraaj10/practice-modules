package service

import (
	"ecommerce-inventory/model"
	"ecommerce-inventory/repository"
	"errors"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// Add a product
func (service *ProductService) AddProduct(product *model.Product) error {
	if product.Name == "" || product.Price <= 0 || product.Stock < 0 {
		return errors.New("invalid product data")
	}
	return service.repo.AddProduct(product)
}

// Get a product by ID
func (service *ProductService) GetProductByID(id int) (*model.Product, error) {
	product, err := service.repo.GetProductByID(id)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

// Update a product
func (service *ProductService) UpdateProduct(product *model.Product) error {
	if product.Name == "" || product.Price <= 0 || product.Stock < 0 {
		return errors.New("invalid product data")
	}
	return service.repo.UpdateProduct(product)
}

// Delete a product
func (service *ProductService) DeleteProduct(id int) error {
	return service.repo.DeleteProduct(id)
}

// Get all products with pagination
func (service *ProductService) GetAllProducts(page, limit int) ([]model.Product, error) {
	return service.repo.GetAllProducts(page, limit)
}
