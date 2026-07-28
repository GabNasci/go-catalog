package usecase

import (
	"go-catalog/model"
	"go-catalog/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repository repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		repository: repository,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}